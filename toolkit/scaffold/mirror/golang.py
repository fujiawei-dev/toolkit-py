from pathlib import Path

import click

from toolkit.config import runtime


@click.command(help="Change golang source mirrors.")
def modify_golang_mirror():
    if runtime.LINUX:
        file = Path.home() / ".config" / "go" / "env"
    elif runtime.WINDOWS:
        file = Path.home() / "AppData" / "Roaming" / "go" / "env"
    else:
        raise RuntimeError("Unsupported platform")

    file.parent.mkdir(parents=True, exist_ok=True)

    content = [
        "GOSUMDB=off",
        "GO111MODULE=on",
        "GOPROXY=https://goproxy.cn,direct",
    ]

    # https://github.com/python/cpython/issues/67894
    file.write_bytes("\n".join(content).encode("utf-8"))

    click.echo("If the setting does not take effect, do it manually:")

    cmd = "go env -w"

    for line in content:
        click.echo(cmd + " " + line)
