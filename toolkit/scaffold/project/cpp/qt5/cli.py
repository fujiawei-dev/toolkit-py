"""Console script for cpp qt5 module."""

import sys

import click

from toolkit.scaffold.project.cpp.qt5.console import create_console
from toolkit.scaffold.project.cpp.qt5.example import create_example
from toolkit.scaffold.project.cpp.qt5.qml import create_qml


@click.group(help="Create a cpp qt5 project scaffold.")
def create_cpp_qt5_project():
    pass


create_cpp_qt5_project.add_command(create_console, "console")

create_cpp_qt5_project.add_command(create_example, "example")

create_cpp_qt5_project.add_command(create_qml, "qml")

if __name__ == "__main__":
    sys.exit(create_cpp_qt5_project())  # pragma: no cover
