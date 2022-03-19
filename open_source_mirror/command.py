"""
Date: 2022.02.02 18:14
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 18:14
"""
import click
from click_aliases import ClickAliasedGroup

from .python import python as _python
from .raspberrypi import Version as raspberrypi_version, raspberrypi as _raspberrypi
from .ubuntu import Version as ubuntu_version, ubuntu as _ubuntu, ubuntu_port


@click.group(cls=ClickAliasedGroup)
def command_cfm():
    pass


@command_cfm.command(
    aliases=["py"],
    context_settings={"help_option_names": ["-h", "--help"]},
    help="Change pypi & conda source minors.",
)
def python():
    _python()


@command_cfm.command(
    aliases=["ub"],
    context_settings={"help_option_names": ["-h", "--help"]},
    help="Change ubuntu/ubuntu-port source minors.",
)
@click.option("--port/--no-port", "-p/", default=False)
@click.option("--version", "-v", default=ubuntu_version.LTS2004)
def ubuntu(port, version):
    if port:
        ubuntu_port(version)
    else:
        _ubuntu(version)


@command_cfm.command(
    aliases=["pi"],
    context_settings={"help_option_names": ["-h", "--help"]},
    help="Change Raspberry Pi OS source minors.",
)
@click.option("--version", "-v", default=raspberrypi_version.Debian10)
def raspberrypi(version):
    _raspberrypi(version)
