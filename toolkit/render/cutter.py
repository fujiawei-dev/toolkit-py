import contextlib
import fnmatch
import json
import os.path
import shutil
import tempfile
from collections import OrderedDict
from pathlib import Path
from typing import Callable, Union

import jinja2
import yaml
from cookiecutter.environment import StrictEnvironment
from cookiecutter.generate import (
    generate_context,
    generate_file,
    is_copy_only_path,
    render_and_create_dir,
)
from cookiecutter.utils import work_in
from jinja2 import FileSystemLoader

from toolkit.config.context import BASE_CONTEXT, COOKIECUTTER_CONTEXT, IGNORED_ITEMS
from toolkit.fs.dir import copy_items_in_directory
from toolkit.logger import logging
from toolkit.template.code_style import get_camel_case_styles

log = logging.getLogger(__name__)


def is_copy_only(path: Union[str, Path], context: dict) -> bool:
    return is_copy_only_path(path, context)


def ignore(path: Union[str, Path], context: dict) -> bool:
    """Check whether the given `path` should only be ignored.

    Returns True if `path` matches a pattern in the given `context` dict,
    otherwise False.

    :param path: A file-system path referring to a file or dir that
        should be rendered or just copied.
    :param context: cookiecutter context.
    """
    ignored = False

    with contextlib.suppress(KeyError):
        for item in context["cookiecutter"]["_ignore"]:
            if fnmatch.fnmatch(path, item):
                ignored = True
                break

    log.debug(f"{path} is ignored: {ignored}")

    return ignored


def generate_cutter_context(
    context_file: str = "cookiecutter.json",
    context: dict = None,
    ignored_items: list = None,
) -> OrderedDict:
    cookiecutter_context = (
        (
            json.load(open(context_file, "r", encoding="utf-8"))
            if os.path.isfile(context_file)
            else {}
        )
        | COOKIECUTTER_CONTEXT
        | {"_ignore": (ignored_items or []) + IGNORED_ITEMS}
    )

    with tempfile.TemporaryDirectory(prefix="context-") as tmpdir:
        temp_context_file = os.path.join(tmpdir, "cookiecutter.json")

        with open(temp_context_file, "w", encoding="utf-8", newline="\n") as f:
            json.dump(cookiecutter_context, f, ensure_ascii=False, indent=4)

        cutter_context = (
            generate_context(temp_context_file) | BASE_CONTEXT | (context or {})
        )

    log.debug(f"Cutter context: {cutter_context}")

    return cutter_context


def generate_rendered_file(
    project_path: str,
    template_file_path: str,
    context: OrderedDict,
    jinja2_environment: jinja2.Environment,
    skip_if_file_exists: bool = True,
):
    return generate_file(
        project_dir=project_path,
        infile=template_file_path,
        context=context,
        env=jinja2_environment,
        skip_if_file_exists=skip_if_file_exists,
    )


def create_rendered_directory(
    project_path: str,
    template_directory_path: str,
    context: OrderedDict,
    jinja2_environment: jinja2.Environment,
    skip_if_directory_exists: bool = True,
):
    return render_and_create_dir(
        dirname=template_directory_path,
        context=context,
        output_dir=project_path,
        environment=jinja2_environment,
        overwrite_if_exists=not skip_if_directory_exists,
    )


def generate_rendered_files_recursively(
    template_path: str,
    project_path: str = ".",
    generated_path_hook: Callable[[str], str] = None,
    user_input_context: dict = None,
    user_input_context_hook: Callable[[dict], None] = None,
    ignored_items: list = None,
    skip_if_file_exists: bool = True,
):
    template_path = os.path.abspath(template_path)
    project_path = os.path.abspath(project_path)

    log.debug("Generating project from %s to %s", template_path, project_path)

    # Generate project files in place instead of creating an extra folder
    project_parent, project_repo = os.path.split(project_path)

    temp_dir = tempfile.mkdtemp()
    temp_project_dir = os.path.join(temp_dir, project_repo)
    os.makedirs(temp_project_dir, exist_ok=True)

    project_slugs = get_camel_case_styles(project_repo)
    context_file = os.path.join(template_path, "cookiecutter.json")
    context = generate_cutter_context(
        context_file=context_file,
        context={
            "project_slug": {
                "words_lowercase": project_slugs[0],  # camel case
                "kebab_case": project_slugs[1],  # camel-case
                "words_capitalized": project_slugs[2],  # Camel Case
                "snake_case": project_slugs[3],  # camel_case
            },
        }
        | (user_input_context or {}),
        ignored_items=ignored_items,
    )

    if callable(user_input_context_hook):
        context = user_input_context_hook(context)

    env_vars = context.get("cookiecutter", {}).get("_jinja2_env_vars", {})

    jinja2_environment = StrictEnvironment(
        context=context,
        keep_trailing_newline=True,
        **env_vars,
    )

    with work_in(template_path):
        jinja2_environment.loader = FileSystemLoader([".", "../templates"])

        for root, dirs, files in os.walk("."):
            copy_dirs = []
            render_dirs = []

            for d in dirs:
                if ignore(d, context):
                    continue

                d_ = os.path.normpath(os.path.join(root, d))
                # We check the full path, because that's how it can be
                # specified in the ``_copy_without_render`` setting, but
                # we store just the dir name
                if is_copy_only(d_, context):
                    log.debug("Found copy only path %s", d)
                    copy_dirs.append(d)
                else:
                    render_dirs.append(d)

            for copy_dir in copy_dirs:
                in_dir = os.path.normpath(os.path.join(root, copy_dir))
                out_dir = os.path.normpath(os.path.join(temp_project_dir, in_dir))
                out_dir = jinja2_environment.from_string(out_dir).render(**context)

                if generated_path_hook:
                    out_dir = generated_path_hook(out_dir)

                log.debug("Copying dir %s to %s without rendering", in_dir, out_dir)

                if os.path.isdir(out_dir):
                    shutil.rmtree(out_dir)

                shutil.copytree(in_dir, out_dir)

            # We mutate ``dirs``, because we only want to go through these dirs recursively
            dirs[:] = render_dirs
            for d in dirs:
                not_rendered_dir = os.path.join(temp_project_dir, root, d)
                create_rendered_directory(
                    project_path=temp_project_dir,
                    template_directory_path=not_rendered_dir,
                    context=context,
                    jinja2_environment=jinja2_environment,
                    skip_if_directory_exists=False,
                )

            for f in files:
                in_file = os.path.normpath(os.path.join(root, f))

                if ignore(in_file, context):
                    continue

                out_file_tmpl = jinja2_environment.from_string(in_file)
                out_file_rendered = out_file_tmpl.render(**context)
                out_file = os.path.join(temp_project_dir, out_file_rendered)

                if is_copy_only_path(in_file, context):
                    log.debug(f"Copying {in_file} to {out_file} without rendering")
                    shutil.copyfile(in_file, out_file)
                    shutil.copymode(in_file, out_file)
                else:
                    generate_rendered_file(
                        project_path=temp_project_dir,
                        template_file_path=in_file,
                        context=context,
                        jinja2_environment=jinja2_environment,
                        skip_if_file_exists=skip_if_file_exists,
                    )

                if generated_path_hook:
                    os.renames(out_file, generated_path_hook(out_file))

    # Move the generated project to the final location
    copy_items_in_directory(temp_project_dir, project_path, skip_if_file_exists)

    # Remove the temp dir
    shutil.rmtree(temp_dir)

    with contextlib.suppress(KeyError):
        del context["cookiecutter"]
        del context["factory"]

    yaml.add_representer(
        OrderedDict,
        lambda dumper, data: dumper.represent_mapping(
            "tag:yaml.org,2002:map", data.items()
        ),
    )

    yaml.dump(
        context,
        open(
            os.path.join(project_path, "cutter.yaml"),
            "w",
            encoding="utf-8",
            newline="\n",
        ),
        sort_keys=True,
        encoding="utf-8",
        allow_unicode=True,
    )
