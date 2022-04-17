"""
Date: 2020-09-21 23:48:26
Description: Change pypi & conda source mirrors.
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 15:01
"""
import os
import subprocess
import sys
import time
from urllib.parse import urlparse

import click


def pypi():
    # https://pypi.org/

    mirrors = [
        "https://pypi.douban.com/simple",
        "https://mirrors.ustc.edu.cn/pypi/web/simple/",
        "https://mirrors.tuna.tsinghua.edu.cn/pypi/web/simple/",
        "https://mirrors.aliyun.com/pypi/simple/",  # very slow
    ]

    cmd = "pip uninstall -y PyQt5-Qt5 && pip install PyQt5-Qt5 --no-cache-dir -i "

    min_cost = 0
    mirror = mirrors[0]

    for index in range(len(mirrors)):
        ts = time.perf_counter()
        subprocess.run(cmd + mirrors[index], shell=True)
        te = time.perf_counter()
        cost = te - ts
        click.echo(f"[{cost:.02f}] {mirrors[index]}")

        if min_cost == 0 or cost < min_cost:
            min_cost = cost
            mirror = mirrors[index]

        if min_cost < 15:
            break

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

    print(conf)


def conda():
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

    print(conf)


def python():
    pypi(), conda()
