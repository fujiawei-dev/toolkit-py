"""Console script for cpp qt6 module."""

import sys

import click

from toolkit.scaffold.project.cpp.qt6.example import create_example


@click.group(help="Create a cpp qt6 project scaffold.")
def create_cpp_qt6_project():
    pass


create_cpp_qt6_project.add_command(create_example, "example")

if __name__ == "__main__":
    sys.exit(create_cpp_qt6_project())  # pragma: no cover
