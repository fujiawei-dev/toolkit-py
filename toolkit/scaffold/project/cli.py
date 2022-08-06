"""Console script for project module."""

import sys

import click

from toolkit.scaffold.project.python.cli import create_python_project


@click.group(help="Create a project scaffold.")
def main():
    pass


main.add_command(create_python_project, "python")

if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
