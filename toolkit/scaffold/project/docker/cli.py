import sys

import click

from toolkit.render.cutter import generate_rendered_files_recursively
from toolkit.scaffold.project import context
from toolkit.scaffold.project.docker.compose import create_compose
from toolkit.scaffold.project.template import TEMPLATE_DOCKER_PATH


@click.group(help="Create a docker project scaffold.")
def create_docker_project():
    pass


@create_docker_project.command(
    name="cli",
    help="Create a docker-cli project scaffold.",
)
@click.option(
    "--project-path",
    type=click.Path(exists=True, file_okay=False),
    default=".",
    help="Project path.",
)
@click.option("--overwrite", is_flag=True, help="Overwrite existing files.")
def create_cli(project_path: str, overwrite: bool):
    user_input_context = context.get_user_input_context()
    ignored_items = context.get_ignored_items({}, [])
    template_path = TEMPLATE_DOCKER_PATH / "cli"

    generate_rendered_files_recursively(
        template_path=template_path,
        project_path=project_path,
        user_input_context=user_input_context,
        ignored_items=ignored_items,
        skip_if_file_exists=not overwrite,
    )


create_docker_project.add_command(create_compose, "compose")

if __name__ == "__main__":
    sys.exit(create_docker_project())  # pragma: no cover
