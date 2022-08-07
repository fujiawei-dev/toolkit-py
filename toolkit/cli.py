"""Console script for toolkit."""
import sys

import click
from click_aliases import ClickAliasedGroup

from toolkit import __version__
from toolkit.provider.image_hosting import ups_for_typora_command
from toolkit.provider.tidy import tidy_command
from toolkit.provider.unzip import unzip_command
from toolkit.provider.user_agent import generate_user_agent_command
from toolkit.provider.youdao import block_youdao_ads_command

__all__ = [
    "generate_user_agent_command",
    "ups_for_typora_command",
]


@click.group(cls=ClickAliasedGroup)
def main():
    pass


@main.command(help="Print the version of toolkit-py.")
def version():
    click.echo(__version__)


main.add_command(tidy_command, name="tidy")

main.add_command(unzip_command, name="unzip")

main.add_command(block_youdao_ads_command, name="youdao")

if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
