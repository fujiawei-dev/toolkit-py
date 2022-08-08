"""Console script for {{project_slug.snake_case}}."""

import sys

import click

from {{project_slug.snake_case}} import __version__


@click.group()
def main():
    pass


@main.command(help="Print the version of {{ project_slug.kebab_case }}.")
def version():
    click.echo(__version__)


if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
