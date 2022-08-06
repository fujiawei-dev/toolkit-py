"""Console script for notes module."""

import sys

import click

from toolkit.scaffold.project.notes.article import create_article


@click.group(help="Create a notes project scaffold or a new article.")
def create_notes_project():
    pass


create_notes_project.add_command(create_article, "article")


@create_notes_project(name="all")
def create_all():
    pass


if __name__ == "__main__":
    sys.exit(create_notes_project())  # pragma: no cover
