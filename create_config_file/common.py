"""
Date: 2022.02.03 9:17
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.03 9:17
"""
import os
import sys

import click


def writer(conf, content="", read_only=True, official=""):
    if not read_only and content != "":
        os.makedirs(os.path.dirname(conf), exist_ok=True)
        with open(conf, "w", encoding="utf8", newline="\n") as fp:
            fp.write(content)

    if official:
        click.echo(official)

    click.echo(conf)


def is_windows():
    return sys.platform.startswith("win")


def join_user(path):
    return os.path.join(os.path.expanduser("~"), path)
