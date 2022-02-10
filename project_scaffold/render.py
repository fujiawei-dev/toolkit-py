"""
Date: 2022.02.10 14:54
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.10 14:54
"""
from pathlib import Path
from typing import Tuple

from jinja2 import Template


def render_by_jinja2(content: str, *args, **kwargs):
    return Template(content).render(*args, **kwargs).strip() + "\n"


def render_by_format(content: str, *args, **kwargs):
    return content.format(*args, **kwargs).strip() + "\n"


def render_templates_recursively(
    p: Path, dst: Path, include_suffixes: Tuple[str], src: Path = None, **kwargs
):
    q = dst / (p.relative_to(src or p))

    if p.is_dir():
        q.mkdir(parents=True, exist_ok=True)
        for e in p.iterdir():
            render_templates_recursively(e, dst, include_suffixes, src or p, **kwargs)

    elif p.is_file():
        suffixes = p.suffixes

        if len(suffixes) > 1:
            suffix = suffixes[0]

            if suffix not in include_suffixes:
                return

            qs = str(q).partition(suffix)
            q = Path(qs[0] + qs[-1])

        kwargs["GOLANG_PACKAGE"] = p.parent.stem
        content = render_by_jinja2(p.open(encoding="utf-8").read(), **kwargs)
        q.open("w", encoding="utf-8", newline="\n").write(content)
