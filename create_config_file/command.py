"""
Date: 2022.02.02 18:14
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 18:14
"""
import click
from click_aliases import ClickAliasedGroup

from .clash import clash as _clash
from .hosts import hosts as _hosts
from .notes import new_article as _new_article
from .python import python as _python


@click.group(cls=ClickAliasedGroup)
def command_ccf():
    pass


@command_ccf.command(
    aliases=["py"],
    context_settings={"help_option_names": ["-h", "--help"]},
    help="Create or display configuration files about Python.",
)
@click.option("--method", "-m", default=1, type=int, help="1 -> .pypirc")
@click.option(
    "--read-only",
    "-r",
    default=True,
    type=bool,
    help="Read only or create configuration files.",
)
def python(method, read_only):
    _python(method, read_only)


@command_ccf.command(help="Display configuration file of hosts.")
def hosts():
    _hosts()


@command_ccf.command(help="Display configuration file of clash.")
def clash():
    _clash()


@command_ccf.command(
    aliases=["nn"],
    context_settings={"help_option_names": ["-h", "--help"]},
    help="Create a new article for hugo.",
)
@click.option(
    "--path", "-p", required=True, type=str, help="The file path for a new article."
)
def notes(path):
    _new_article(path)
