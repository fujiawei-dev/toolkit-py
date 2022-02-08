"""
Date: 2022.02.07 15:08
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.07 15:08
"""
import os

import click
from click_aliases import ClickAliasedGroup

from .auto_unzip import PASSWORDS_CUSTOMIZE_FILE, PASSWORDS_DEFAULT_DIR, Unzipper


@click.group(cls=ClickAliasedGroup)
def command_ucmd():
    pass


@command_ucmd.command(
    aliases=["unzip", "uz"],
    context_settings={"help_option_names": ["-h", "--help"]},
    help="Automatically unzip files recursively.",
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
def auto_unzip(config, test):
    unzipper = Unzipper()

    if config:
        click.echo(PASSWORDS_DEFAULT_DIR + "\n" + PASSWORDS_CUSTOMIZE_FILE)
        return

    if test:
        click.echo(unzipper.create_7z_files_for_test())
        return

    unzipper.run(os.getcwd())
