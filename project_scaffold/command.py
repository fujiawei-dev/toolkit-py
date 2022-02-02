'''
Date: 2022.02.02 18:14
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 18:14
'''
import click

from .python import python as _python


@click.group()
def command_cps():
    pass


@command_cps.command(help="Create python project scaffold.")
def py():
    _python()


@command_cps.command(help="Create python project scaffold.")
def python():
    _python()
