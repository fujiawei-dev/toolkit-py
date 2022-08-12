"""Console script for golang module."""

import sys

import click

from toolkit.scaffold.project.golang.example import create_example
from toolkit.scaffold.project.golang.factory import create_factory
from toolkit.scaffold.project.golang.web import create_web


@click.group(help="Create a golang project scaffold.")
def create_golang_project():
    pass


create_golang_project.add_command(create_example, "example")

create_golang_project.add_command(create_factory, "factory")

create_golang_project.add_command(create_web, "web")

if __name__ == "__main__":
    sys.exit(create_golang_project())  # pragma: no cover
