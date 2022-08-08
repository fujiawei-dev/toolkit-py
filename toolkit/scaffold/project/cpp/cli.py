"""Console script for cpp module."""

import sys

import click

from toolkit.scaffold.project.cpp.qt5.cli import create_cpp_qt5_project


@click.group(help="Create a cpp project scaffold.")
def create_cpp_project():
    pass


create_cpp_project.add_command(create_cpp_qt5_project, "qt5")

if __name__ == "__main__":
    sys.exit(create_cpp_project())  # pragma: no cover
