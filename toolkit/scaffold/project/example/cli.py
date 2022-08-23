"""Console script for example module."""

import sys

import click

from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.template import TEMPLATE_EXAMPLE_PATH

create_all = generate_create_project_command(
    command_help="Create a example project scaffold.",
    template_paths=TEMPLATE_EXAMPLE_PATH,
)


@click.group(
    help="Create a example project scaffold.",
    invoke_without_command=True,
)
@click.pass_context
def create_example_project(ctx: click.Context):
    if ctx.invoked_subcommand is None:
        ctx.invoke(create_all)


create_example_project.add_command(create_all, "all")

if __name__ == "__main__":
    sys.exit(create_example_project())  # pragma: no cover
