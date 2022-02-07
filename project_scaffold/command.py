"""
Date: 2022.02.02 18:14
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 18:14
"""
import click

from .c import c as _c
from .common import create_common_files
from .notes import notes as _notes
from .python import python as _python


@click.group()
def command_cps():
    pass


@command_cps.command(help="Create basic project scaffold.")
def base():
    create_common_files()


@command_cps.command(help="Create Python project scaffold.")
def py():
    _python()


@command_cps.command(help="Create Python project scaffold.")
def python():
    _python()


@command_cps.command(help="Create notes project scaffold.")
def notes():
    _notes()


@command_cps.command(help="Create C/C++ project scaffold.")
def c():
    _c()
