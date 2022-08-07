import click

from toolkit.scaffold.project.render import generate_rendered_project


def generate_create_project_command(
    command_help: str,
    template_path: str,
    raw_context: dict = None,
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
            raw_context=raw_context,
            project_context=project_context,
            ignored_fields=ignored_fields,
            overwrite=overwrite,
        )

    return create_project
