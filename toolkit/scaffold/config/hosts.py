import click

from toolkit.config.runtime import EDITOR, WINDOWS


@click.command(help="Edit hosts file.")
def edit_hosts():
    filename = "C:/Windows/System32/drivers/etc/hosts" if WINDOWS else "/etc/hosts"
    click.edit(filename=filename, editor=EDITOR)
