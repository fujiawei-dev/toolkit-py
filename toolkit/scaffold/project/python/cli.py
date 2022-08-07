"""Console script for python module."""

import sys

import click

from toolkit.scaffold.project.python.package import create_package


@click.group(help="Create a python project scaffold.")
def create_python_project():
    pass


create_python_project.add_command(create_package, "package")

if __name__ == "__main__":
    sys.exit(create_python_project())  # pragma: no cover
