import click
from click_aliases import ClickAliasedGroup

from toolkit import __version__


@click.group(cls=ClickAliasedGroup)
def main():
    pass


@main.command(help="Print the version of toolkit-py.")
def version():
    click.echo(__version__)
