import ctypes
import sys
from pathlib import Path

import click

INSTALL_DIR = Path("C:/Program Files (x86)/Youdao/YoudaoNote")


def block_ads():
    if (file := INSTALL_DIR / "theme" / "build.xml").exists():
        file.write_text(file.read_text().replace("161", "0").replace("250,160", "0,0"))


def is_admin():
    return ctypes.windll.shell32.IsUserAnAdmin()


def block_ads_as_admin():
    if is_admin():
        block_ads()
    else:
        ctypes.windll.shell32.ShellExecuteW(
            None, "runas", sys.executable, __file__, None, 1
        )


@click.command(help="Remove the ui ads of YoudaoNote (run as admin)")
def block_youdao_ads_command():
    if sys.platform != "win32":
        exit(0)

    block_ads()


if __name__ == "__main__":
    block_ads_as_admin()
