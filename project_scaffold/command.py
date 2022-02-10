"""
Date: 2022.02.02 18:14
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 18:14
"""
import click
from click_aliases import ClickAliasedGroup

from .c import c as _c
from .common import create_common_files
from .golang import golang as _golang
from .notes import notes as _notes
from .python import python as _python


@click.group(cls=ClickAliasedGroup)
def command_cps():
    pass


@command_cps.command(help="Create basic project scaffold.")
def base():
    create_common_files()


@command_cps.command(
    aliases=["py"],
    context_settings={"help_option_names": ["-h", "--help"]},
    help="Create Python project scaffold.",
)
def python():
    _python()


@command_cps.command(
    aliases=["go"],
    context_settings={"help_option_names": ["-h", "--help"]},
    help="Create Golang project scaffold.",
)
def golang():
    _golang()


@command_cps.command(help="Create notes project scaffold.")
def notes():
    _notes()


@command_cps.command(help="Create C/C++ project scaffold.")
def c():
    _c()
