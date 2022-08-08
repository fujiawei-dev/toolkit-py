from typing import Callable

import click

from toolkit.scaffold.project.render import generate_rendered_project


def generate_create_project_command(
    command_help: str,
    template_path: str,
    generated_path_hook: Callable[[str], str] = None,
    raw_user_input_context: dict = None,
    user_input_context_hook: Callable[[dict], dict] = None,
    project_context: dict = None,
    ignored_fields: list = None,
) -> click.Command:
    @click.command(
        help=command_help,
        context_settings={"ignore_unknown_options": True},
    )
    @click.option(
        "--project-path",
        type=click.Path(exists=True, file_okay=False),
        default=".",
        help="The path to the final generated project.",
    )
    @click.option(
        "--overwrite",
        "-y",
        is_flag=True,
        help="Overwrite existing files.",
    )
    def create_project(project_path: str, overwrite: bool):
        generate_rendered_project(
            template_path=template_path,
            project_path=project_path,
            generated_path_hook=generated_path_hook,
            raw_user_input_context=raw_user_input_context,
            user_input_context_hook=user_input_context_hook,
            project_context=project_context,
            ignored_fields=ignored_fields,
            overwrite=overwrite,
        )

    return create_project
