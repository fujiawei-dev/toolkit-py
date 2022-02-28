"""
Date: 2022.02.03 9:17
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.03 9:17
"""
import os
import sys

import click

SEPARATE = "--------------\n"


def writer(*confs, content="", read_only=True, official="", message="", append=False):
    if not read_only and content != "":
        for conf in confs:
            os.makedirs(os.path.dirname(conf), exist_ok=True)
            conf_backup = conf + ".backup"
            if os.path.exists(conf) and not os.path.exists(conf_backup):
                os.renames(conf, conf_backup)
            mode = "a" if append else "w"
            with open(conf, mode, encoding="utf-8", newline="\n") as fp:
                fp.write(content.strip() + "\n")
            click.echo(f"written to {conf}, and backup to {conf_backup}")
    else:
        click.echo(f"read only {', '.join(confs)}")

    if official:
        click.echo(official)

    if message:
        click.echo(message)


def is_windows():
    return sys.platform.startswith("win")


def join_user(path):
    return os.path.join(os.path.expanduser("~"), path)
