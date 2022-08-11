from pathlib import Path
from typing import Callable, Union

import click

from toolkit.scaffold.project.render import generate_rendered_project


def generate_create_project_command(
    command_help: str,
    template_paths: Union[Union[str, Path], list[Union[str, Path]]],
    generated_path_hook: Callable[[str], str] = None,
    raw_user_input_context: dict = None,
    factory_user_input_context: dict = None,
    user_input_context_hook: Callable[[dict], dict] = None,
    project_context: dict = None,
    ignored_fields: list = None,
) -> click.Command:
    """
    Generate a click command that can be used to create a project scaffold.

    Args:
        command_help: The help message for the command.
        template_paths: The paths to the templates to use for the project scaffold.
        generated_path_hook: A function that can be used to modify the generated path.
        raw_user_input_context: The raw user input context to use for the project scaffold.
        factory_user_input_context: The user input context to use for the project scaffold.
        user_input_context_hook: A function that can be used to modify the user input context.
        project_context: The project context to use for the project scaffold.
        ignored_fields: The fields to ignore when generating the project scaffold.

    Returns:
        The click command.
    """
    if isinstance(template_paths, (str, Path)):
        template_paths = [template_paths]

    assert len(template_paths) > 0, "No template paths provided."

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
            template_paths=template_paths,
            project_path=project_path,
            generated_path_hook=generated_path_hook,
            raw_user_input_context=raw_user_input_context,
            factory_user_input_context=factory_user_input_context,
            user_input_context_hook=user_input_context_hook,
            project_context=project_context,
            ignored_fields=ignored_fields,
            overwrite=overwrite,
        )

    return create_project
