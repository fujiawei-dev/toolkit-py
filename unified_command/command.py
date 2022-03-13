"""
Date: 2022.02.07 15:08
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.07 15:08
"""
import os

import click
from click_aliases import ClickAliasedGroup

from .auto_unzip import PASSWORDS_DIR, PASSWORDS_FILE, Unzipper
from .version import __version__
from .youdao import remove_youdao_note_ad


@click.group(cls=ClickAliasedGroup)
def command_ucmd():
    pass


@command_ucmd.command()
def version():
    click.echo(__version__)


@command_ucmd.command(
    aliases=["unzip", "uz"],
    context_settings={"help_option_names": ["-h", "--help"]},
    help="Automatically unzip files recursively.",
)
@click.option(
    "--path",
    "-p",
    type=str,
    help="The source path.",
)
@click.option(
    "--config",
    "-c",
    is_flag=True,
    help="Show the default path of configuration file.",
)
@click.option(
    "--test",
    "-t",
    is_flag=True,
    help="Create 7z files for test.",
)
def auto_unzip(path, config, test):
    unzipper = Unzipper()

    if config:
        click.echo(PASSWORDS_DIR + "\n" + PASSWORDS_FILE)
        return

    if test:
        click.echo(unzipper.create_7z_files_for_test())
        return

    unzipper.run_with_history(path if path and os.path.exists(path) else os.getcwd())


@command_ucmd.command(help="Remove the ui ads of YoudaoNote")
def youdao():
    remove_youdao_note_ad()
