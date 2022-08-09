"""Console script for notes module."""

import sys

import click

from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.notes.article import create_article
from toolkit.scaffold.project.notes.context import NOTES_USER_INPUT_CONTEXT
from toolkit.scaffold.project.template import TEMPLATE_NOTES_PATH

create_all = generate_create_project_command(
    command_help="Create a notes project scaffold.",
    template_paths=TEMPLATE_NOTES_PATH,
    raw_user_input_context=NOTES_USER_INPUT_CONTEXT,
)


@click.group(
    help="Create a notes project scaffold or a new article.",
    invoke_without_command=True,
)
@click.pass_context
def create_notes_project(ctx: click.Context):
    if ctx.invoked_subcommand is None:
        ctx.invoke(create_all)


create_notes_project.add_command(create_article, "article")

create_notes_project.add_command(create_all, "all")

if __name__ == "__main__":
    sys.exit(create_notes_project())  # pragma: no cover
