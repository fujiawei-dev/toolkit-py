"""
Date: 2022.02.02 18:14
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 18:14
"""
import click

from .python import python as _python


@click.group()
def command_cfm():
    pass


@command_cfm.command(help="Change pypi & conda source minors.")
def py():
    _python()


@command_cfm.command(help="Change pypi & conda source minors.")
def python():
    _python()
