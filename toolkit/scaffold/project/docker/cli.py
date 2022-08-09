import sys

import click

from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.docker.compose import create_compose
from toolkit.scaffold.project.template import TEMPLATE_DOCKER_PATH


@click.group(help="Create a docker project scaffold.")
def create_docker_project():
    pass


create_cli = generate_create_project_command(
    command_help="Create a docker-cli project scaffold.",
    template_paths=TEMPLATE_DOCKER_PATH / "cli",
)

create_docker_project.add_command(create_cli, "cli")

create_docker_project.add_command(create_compose, "compose")

if __name__ == "__main__":
    sys.exit(create_docker_project())  # pragma: no cover
