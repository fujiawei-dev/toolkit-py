import os.path
from pathlib import Path
from typing import Callable, Union

import yaml

from toolkit.render.cutter import generate_rendered_files_recursively
from toolkit.scaffold.project import context


def generate_rendered_project(
    template_paths: Union[Union[str, Path], list[Union[str, Path]]],
    project_path: str = ".",
    generated_path_hook: Callable[[str], str] = None,
    raw_user_input_context: dict = None,
    factory_user_input_context: dict = None,
    user_input_context_hook: Callable[[dict], dict] = None,
    project_context: dict = None,
    ignored_fields: list = None,
    overwrite: bool = False,
):
    if not os.path.exists(project_path):
        os.makedirs(project_path, exist_ok=True)

    if isinstance(template_paths, (str, Path)):
        template_paths = [template_paths]

    assert len(template_paths) > 0, "No template paths provided."

    cutter_file = os.path.join(project_path, "cutter.yaml")

    if os.path.exists(cutter_file):
        user_input_context = yaml.safe_load(open(cutter_file))
    else:
        user_input_context = context.get_user_input_context(raw_user_input_context)

    if factory_user_input_context:
        factory = context.get_user_input_context(factory_user_input_context)
        user_input_context["factory"] = factory

    ignored_items = context.get_ignored_items(project_context, ignored_fields)

    for i in range(len(template_paths)):
        generate_rendered_files_recursively(
            template_path=template_paths[i],
            project_path=project_path,
            generated_path_hook=generated_path_hook,
            user_input_context=user_input_context,
            user_input_context_hook=user_input_context_hook,
            ignored_items=ignored_items,
            skip_if_file_exists=not overwrite and i == 0,
        )
