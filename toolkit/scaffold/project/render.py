from typing import Callable

from toolkit.render.cutter import generate_rendered_files_recursively
from toolkit.scaffold.project import context


def generate_rendered_project(
    template_path: str,
    project_path: str = ".",
    generated_path_hook: Callable[[str], str] = None,
    raw_user_input_context: dict = None,
    user_input_context_hook: Callable[[dict], dict] = None,
    project_context: dict = None,
    ignored_fields: list = None,
    overwrite: bool = False,
):
    user_input_context = context.get_user_input_context(raw_user_input_context)

    ignored_items = context.get_ignored_items(project_context, ignored_fields)

    generate_rendered_files_recursively(
        template_path=template_path,
        project_path=project_path,
        generated_path_hook=generated_path_hook,
        user_input_context=user_input_context,
        user_input_context_hook=user_input_context_hook,
        ignored_items=ignored_items,
        skip_if_file_exists=not overwrite,
    )
