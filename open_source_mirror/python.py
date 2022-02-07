"""
Date: 2020-09-21 23:48:26
Description: Change pypi & conda source minors.
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 15:01
"""
import os
import sys


def pypi():
    # https://pypi.org/
    # https://mirrors.ustc.edu.cn/pypi/web/simple/
    # https://mirrors.tuna.tsinghua.edu.cn/pypi/web/simple/
    # https://mirrors.aliyun.com/pypi/simple/

    user = os.path.expanduser("~")
    conf = os.path.join(user, ".pip/pip.conf")

    if sys.platform.startswith("win"):
        conf = os.path.join(user, "AppData/Roaming/pip/pip.ini")

    os.makedirs(os.path.dirname(conf), exist_ok=True)

    with open(conf, "w", encoding="utf-8", newline="\n") as fp:
        fp.write(
            "[global]\n"
            "index-url=https://mirrors.aliyun.com/pypi/simple/\n"
            "[install]\n"
            "trusted-host=mirrors.aliyun.com\n",
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
