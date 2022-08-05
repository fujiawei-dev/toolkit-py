"""Change pypi & conda source mirrors."""

import os
import subprocess
import sys
import time
from urllib.parse import urlparse

import click


def modify_pypi_mirror():
    # https://pypi.org/

    mirrors = [
        "https://pypi.douban.com/simple",
        "https://mirrors.ustc.edu.cn/pypi/web/simple/",
        "https://mirrors.tuna.tsinghua.edu.cn/pypi/web/simple/",
        "https://mirrors.aliyun.com/pypi/simple/",  # very slow
    ]

    packages = ["httpie", "scapy"]

    uninstall_args = ["pip", "uninstall", "-y", *packages]

    install_args = ["pip", "install", *packages, "--no-cache-dir", "-i"]

    min_cost = 0
    mirror = mirrors[0]

    for index in range(len(mirrors)):
        subprocess.check_call(uninstall_args)

        ts = time.perf_counter()
        subprocess.check_call([*install_args, mirrors[index]])
        te = time.perf_counter()

        cost = te - ts
        click.echo(f"[{cost:.02f}] {mirrors[index]}")

        if min_cost == 0 or cost < min_cost:
            min_cost = cost
            mirror = mirrors[index]

        if min_cost < 10:
            break

    click.echo(f"[faster] {mirror}")

    user = os.path.expanduser("~")
    conf = os.path.join(user, ".pip/pip.conf")

    if sys.platform.startswith("win"):
        conf = os.path.join(user, "AppData/Roaming/pip/pip.ini")

    os.makedirs(os.path.dirname(conf), exist_ok=True)

    with open(conf, "w", encoding="utf-8", newline="\n") as fp:
        fp.write(
            "[global]\n"
            f"index-url={mirror}\n"
            "[install]\n"
            f"trusted-host={urlparse(mirror).netloc}\n",
        )


def modify_conda_mirror():
    # https://www.anaconda.com/

    conf = os.path.join(os.path.expanduser("~"), ".condarc")

    with open(conf, "w", encoding="utf-8", newline="\n") as fp:
        fp.write(
            "channels:\n"
            "  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main/\n"
            "  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/\n"
            "  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/conda-forge/\n"
            "  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/pytorch/\n"
            "  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/pytorch3d/\n"
            "ssl_verify: false\n"
            "auto_activate_base: false\n"
        )


@click.command(help="Change pypi & conda mirrors.")
def modify_python_mirror():
    modify_pypi_mirror(), modify_conda_mirror()
