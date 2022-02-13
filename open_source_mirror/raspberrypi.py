"""
Date: 2022.02.13 14:20
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.13 14:20
"""
from enum import Enum

from create_config_file.common import writer


class Version(str, Enum):
    Debian09 = "stretch"
    Debian10 = "buster"
    Debian11 = "bullseye"


def raspberrypi(version=Version.Debian10):
    """x86 amd64"""
    writer(
        "/etc/apt/sources.list",
        f"""deb [arch=armhf] http://mirrors.tuna.tsinghua.edu.cn/raspbian/raspbian/ {version} main non-free contrib rpi
        deb-src http://mirrors.tuna.tsinghua.edu.cn/raspbian/raspbian/ {version} main non-free contrib rpi
        deb [arch=arm64] http://mirrors.tuna.tsinghua.edu.cn/raspbian/multiarch/ {version} main""",
        read_only=False,
        official="https://mirrors.tuna.tsinghua.edu.cn/help/raspbian/",
    )

    writer(
        "/etc/apt/sources.list.d/raspi.list",
        "deb http://mirrors.tuna.tsinghua.edu.cn/raspberrypi/ {version} main",
        read_only=False,
        official="https://mirrors.tuna.tsinghua.edu.cn/help/raspbian/",
    )
