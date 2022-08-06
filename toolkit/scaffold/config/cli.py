"""Console script for config module."""

import sys

import click
from click_aliases import ClickAliasedGroup

from toolkit.scaffold.config.alias import print_custom_aliases
from toolkit.scaffold.config.git import create_or_edit_git_config
from toolkit.scaffold.config.hosts import edit_hosts
from toolkit.scaffold.config.pwsh import create_or_edit_powershell_profile
from toolkit.scaffold.config.python import create_or_edit_python_config
from toolkit.scaffold.config.vim import create_or_edit_vim_config


@click.group(cls=ClickAliasedGroup)
def main():
    pass


main.add_command(create_or_edit_git_config, "git")

main.add_command(create_or_edit_powershell_profile, "pwsh")

main.add_command(create_or_edit_python_config, "python")

main.add_command(create_or_edit_vim_config, "vim")

main.add_command(edit_hosts, "hosts")

main.add_command(print_custom_aliases, "alias")


if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
