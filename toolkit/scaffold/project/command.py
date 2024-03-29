import os.path
from pathlib import Path
from typing import Callable, Union

import click
from cookiecutter.prompt import read_user_choice, read_user_yes_no

from toolkit.scaffold.project.render import generate_rendered_project


def generate_create_project_command(
    command_help: str,
    template_paths: Union[Union[str, Path], list[Union[str, Path]]],
    generated_path_hook: Callable[[str], str] = None,
    raw_user_input_context: dict = None,
    factory_user_input_context: dict = None,
    factory_user_input_context_hook: Callable[[dict], dict] = None,
    user_input_context_hook: Callable[[dict], dict] = None,
    project_context: dict = None,
    ignored_fields: list = None,
    editors: Union[list, str] = None,
) -> click.Command:
    """
    Generate a click command that can be used to create a project scaffold.

    Args:
        command_help: The help message for the command.
        template_paths: The paths to the templates to use for the project scaffold.
        generated_path_hook: A function that can be used to modify the generated path.
        raw_user_input_context: The raw user input context to use for the project scaffold.
        factory_user_input_context: The user input context to use for the project scaffold.
        factory_user_input_context_hook: A function that can be used to modify the factory user input context.
        user_input_context_hook: A function that can be used to modify the user input context.
        project_context: The project context to use for the project scaffold.
        ignored_fields: The fields to ignore when generating the project scaffold.
        editors: The editors to use when generating the project scaffold.

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
        "--ignored-items",
        type=click.STRING,
        default="",
        help="The items to ignore when generating the project scaffold.",
    )
    @click.option(
        "--launch-editor",
        "-e",
        is_flag=True,
        help="Launch the editor after generating the project.",
    )
    @click.option("--overwrite", "-y", is_flag=True, help="Overwrite existing files.")
    @click.argument("project-path", type=click.Path(file_okay=False), default=".")
    def create_project(
        ignored_items: str,
        launch_editor: bool,
        overwrite: bool,
        project_path: str,
    ):
        """
        Create a project scaffold.

        Args:
            ignored_items: The items to ignore when generating the project scaffold.
            launch_editor: Whether to launch the editor after generating the project.
            overwrite: Whether to overwrite existing files.
            project_path: The path to the project to create.

        Examples:
            Create a project scaffold:

            $ project <language> <type> <project-path>
        """
        generate_rendered_project(
            template_paths=template_paths,
            project_path=project_path,
            generated_path_hook=generated_path_hook,
            raw_user_input_context=raw_user_input_context,
            factory_user_input_context=factory_user_input_context,
            factory_user_input_context_hook=factory_user_input_context_hook,
            user_input_context_hook=user_input_context_hook,
            project_context=project_context,
            ignored_fields=ignored_fields,
            ignored_items=ignored_items.split(","),
            overwrite=overwrite,
        )

        click.echo(f"Project created at:\n{os.path.abspath(project_path)}")

        if project_path != "." and (
            launch_editor or read_user_yes_no("Launch editor?", "yes")
        ):
            click.edit(
                editor=editors
                if isinstance(editors, str)
                else (
                    editors[0]
                    if len(editors) == 1
                    else read_user_choice(
                        "Which editor to use?",
                        editors
                        or [
                            "code",
                            "clion",
                            "goland",
                            "pycharm",
                        ],
                    )
                ),
                filename=project_path,
                require_save=False,
            )

    return create_project
