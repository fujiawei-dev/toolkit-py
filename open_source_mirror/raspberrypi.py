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


def raspberrypi(version=Version.Debian11):
    sources = open("/etc/apt/sources.list").read()

    for version in Version:
        if version in sources:
            break

    print(f"version: {version}")

    writer(
        "/etc/apt/sources.list",
        content=f"""deb [arch=armhf] http://mirrors.tuna.tsinghua.edu.cn/raspbian/raspbian/ {version} main non-free contrib rpi
        deb-src http://mirrors.tuna.tsinghua.edu.cn/raspbian/raspbian/ {version} main non-free contrib rpi
        deb [arch=arm64] http://mirrors.tuna.tsinghua.edu.cn/raspbian/multiarch/ {version} main""",
        read_only=False,
    )

    writer(
        "/etc/apt/sources.list.d/raspi.list",
        content=f"deb http://mirrors.tuna.tsinghua.edu.cn/raspberrypi/ {version} main",
        read_only=False,
        official="https://mirrors.tuna.tsinghua.edu.cn/help/raspbian/",
    )
