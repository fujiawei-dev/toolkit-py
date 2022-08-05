"""Console script for mirror module."""

import sys

import click
from click_aliases import ClickAliasedGroup

from toolkit.scaffold.mirror.golang import modify_golang_mirror
from toolkit.scaffold.mirror.python import modify_python_mirror


@click.group(cls=ClickAliasedGroup)
def main():
    pass


main.add_command(modify_golang_mirror, "go")
main.add_command(modify_python_mirror, "python")


if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
