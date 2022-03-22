"""
Date: 2022.02.10 14:54
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.10 14:54
"""
import os.path
import time
from pathlib import Path
from typing import List, Tuple

import click
from jinja2 import Template

from .common import (
    GENERATOR_HEADER,
    create_common_files,
    get_different_camel_case_styles,
)
from .templates import TEMPLATES_PATH


def render_by_jinja2(content: str, *args, **kwargs):
    return Template(content).render(*args, **kwargs).strip() + "\n"


def render_by_format(content: str, *args, **kwargs):
    return (
        content.strip()
        .replace("{{", "{{{{")
        .replace("}}", "}}}}")
        .format(*args, **kwargs)
        + "\n"
    )


def render_templates_recursively(
    p: Path,
    dst: Path,
    special_paths: Tuple[str] = None,
    include_suffixes: Tuple[str] = None,
    src: Path = None,
    render=render_by_jinja2,
    only_files: Tuple[str] = None,
    replace_list: dict = None,
    **kwargs,
):
    src = src or p
    q = dst / (p.relative_to(src))

    if p.is_dir():
        if len(p.suffixes) > 0:
            q = q.parent.joinpath(*render(q.name, **kwargs).strip().partition("."))
        q.mkdir(parents=True, exist_ok=True)

        for e in p.iterdir():
            render_templates_recursively(
                e,
                q,
                special_paths=special_paths,
                render=render_by_format
                if render != render_by_format
                and special_paths
                and any(
                    map(
                        lambda op: os.path.normpath(p) == os.path.normpath(src / op),
                        special_paths,
                    )
                )
                else render,
                include_suffixes=include_suffixes,
                src=p,
                only_files=only_files,
                replace_list=replace_list,
                **kwargs,
            )

    elif p.is_file():
        if only_files and (
            p.stem.lower() not in only_files
            and p.stem.lower().partition(".")[0] not in only_files
            and p.name.lower() not in only_files
        ):
            return

        suffixes = p.suffixes

        if len(suffixes) > 1:
            suffix = suffixes[-2]
            last_suffix = suffixes[-1]

            if last_suffix == ".py":
                qs = render(q.name, **kwargs).strip().partition(".")
                q = q.parent / qs[0] / qs[-1]
                q.parent.mkdir(parents=True, exist_ok=True)

            elif include_suffixes and suffix in include_suffixes:
                qs = str(q).partition(suffix)
                q = Path(qs[0] + qs[-1])
            else:
                return

        if replace_list:
            s = q.stem
            for k, v in replace_list.items():
                s = s.replace(k, v)
            q = q.with_stem(s)

        if q.exists():
            return

        kwargs["GOLANG_PACKAGE"] = p.parent.stem

        try:
            content = render(p.open(encoding="utf-8").read(), **kwargs)
            if replace_list:
                for k, v in replace_list.items():
                    content = content.replace(k, v)
            q.open("w", encoding="utf-8", newline="\n").write(content)
        except UnicodeDecodeError:
            q.write_bytes(p.read_bytes())

        click.echo(f"created: {q.relative_to(Path.cwd())}")


def render_templates(
    relpath,
    special_paths: List[str] = None,
    include_suffixes: List[str] = None,
    folders: List[str] = None,
    common: bool = True,
    only_files: str = "",
    replace_list: dict = None,
    **kwargs,
):
    package, package_title, package_underscore = get_different_camel_case_styles()

    render_templates_recursively(
        TEMPLATES_PATH / relpath,
        Path.cwd(),
        special_paths=special_paths,
        include_suffixes=include_suffixes,
        only_files=only_files.lower().split(";") if only_files else None,
        replace_list=replace_list,
        **{
            "STUDY_OBJECT": package_title,
            "PACKAGE_TITLE": package_title,
            "APP_NAME": package_underscore,
            "APP_NAME_UPPER": package_underscore.upper(),
            "GOLANG_MODULE": package,
            "PYPI_PACKAGE": package,
            "PYTHON_MODULE": package_underscore,
            "HASHTAG_COMMENTS": GENERATOR_HEADER,
            "SLASH_COMMENTS": GENERATOR_HEADER.replace("#", "//"),
            "CREATED_AT": time.strftime("%Y-%m-%dT%H:%M:%S+08:00"),
            **kwargs,
        },
    )

    if common:
        create_common_files(folders)
    else:
        for folder in folders:
            os.makedirs(folder, exist_ok=True)
