"""
Date: 2022.02.02 18:14
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 18:14
"""
import os
import shutil

import click
from click_aliases import ClickAliasedGroup

from .c import Qt5Templates, c as _c, qt5 as _qt5
from .common import create_common_files
from .golang import GoCombinations, golang as _golang
from .notes import notes as _notes
from .python import python as _python


@click.group(cls=ClickAliasedGroup)
def command_cps():
    pass


@command_cps.command(help="Create basic project scaffold.")
def base():
    create_common_files()


@command_cps.command(
    aliases=["rm", "clean", "clear", "release"],
    help="Remove all example files for release.",
)
def clean():
    build_dirs = {
        "cmake-build-debug",  # Clion cmake build
        "cmake-build-release",  # Makefile cmake build
    }

    for root, dirs, files in os.walk("."):
        for f in files:
            if f.find("example") != -1 or f in {"ToolkitPy_logo.png"}:
                f = os.path.join(root, f)
                click.echo(f"clean {f}")
                os.unlink(f)

        for d in dirs:
            if d.find("example") != -1:
                d = os.path.join(root, d)
                click.echo(f"clean {d}")
                os.rmdir(d)

            elif d in build_dirs:
                d = os.path.join(root, d)
                click.echo(f"clean {d}")
                shutil.rmtree(d)
                os.mkdir(d)


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
@click.option(
    "--combination",
    "-c",
    type=click.Choice(GoCombinations),
    help="Combination of frameworks.",
)
@click.option(
    "--entity", "-e", type=str, default="", help="New entity instance.",
)
def golang(combination, entity):
    _golang(combination, entity)


@command_cps.command(help="Create notes project scaffold.")
def notes():
    _notes()


@command_cps.command(help="Create C/C++ project scaffold.")
@click.option(
    "--only-files",
    "-o",
    type=str,
    default="",
    help="Only create the specified C/C++ template files, split by semicolon.",
)
def c(only_files):
    _c(only_files)


@command_cps.command(help="Create Qt5 project scaffold.")
@click.option(
    "--template",
    "-t",
    type=click.Choice(Qt5Templates),
    default=Qt5Templates.Gui,
    help="Type of Qt5 template files.",
)
@click.option(
    "--only-files",
    "-o",
    type=str,
    default="",
    help="Only create the specified template files, split by semicolon.",
)
def qt5(template=Qt5Templates.Gui, only_files=""):
    _qt5(template, only_files)
