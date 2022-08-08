"""Console script for cpp qt5 module."""

import sys

import click

from toolkit.scaffold.project.cpp.qt5.console import create_console


@click.group(help="Create a cpp qt5 project scaffold.")
def create_cpp_qt5_project():
    pass


create_cpp_qt5_project.add_command(create_console, "console")

if __name__ == "__main__":
    sys.exit(create_cpp_qt5_project())  # pragma: no cover
