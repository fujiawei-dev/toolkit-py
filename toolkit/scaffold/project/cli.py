"""Console script for project module."""

import sys

import click

from toolkit.scaffold.project.ansible.cli import create_ansible_project
from toolkit.scaffold.project.cpp.cli import create_cpp_project
from toolkit.scaffold.project.docker.cli import create_docker_project
from toolkit.scaffold.project.golang.cli import create_golang_project
from toolkit.scaffold.project.notes.cli import create_notes_project
from toolkit.scaffold.project.python.cli import create_python_project


@click.group(help="Create a project scaffold.")
def main():
    pass


main.add_command(create_ansible_project, "ansible")

main.add_command(create_cpp_project, "cpp")

main.add_command(create_docker_project, "docker")

main.add_command(create_golang_project, "golang")

main.add_command(create_notes_project, "notes")

main.add_command(create_python_project, "python")

if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
