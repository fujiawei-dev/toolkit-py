'''
Date: 2022.02.02 18:14
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 18:14
'''
import click
from click_aliases import ClickAliasedGroup

from .python import python as _python


@click.group(cls=ClickAliasedGroup)
def command_ccf():
    pass


@command_ccf.command(
        aliases=['py'],
        context_settings={'help_option_names': ['-h', '--help']},
        help="Create or display configuration files about Python."
)
@click.option('--method', '-m', default=1, type=int, help="1 -> .pypirc")
@click.option('--read-only', '-r', default=True, type=bool, help="Read only or create configuration files.")
def python(method, read_only):
    _python(method, read_only)
