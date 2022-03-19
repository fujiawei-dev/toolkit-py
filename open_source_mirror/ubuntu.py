"""
Date: 2022.02.13 13:41
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.13 13:41
"""
from enum import Enum

from create_config_file.common import writer


class Version(str, Enum):
    LTS1604 = "xenial"
    LTS1804 = "bionic"
    LTS2004 = "focal"


def ubuntu(version=Version.LTS2004):
    """x86 amd64"""
    sources = open("/etc/apt/sources.list").read()

    for version in Version:
        if version in sources:
            break

    print(f"version: {version}")

    content = f"""deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ {version} main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ {version} main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ {version}-updates main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ {version}-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ {version}-backports main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ {version}-backports main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ {version}-security main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ {version}-security main restricted universe multiverse"""

    conf = "/etc/apt/sources.list"
    writer(
        conf,
        content=content,
        read_only=False,
        official="https://mirrors.tuna.tsinghua.edu.cn/help/ubuntu/",
    )


def ubuntu_port(version=Version.LTS2004):
    """arm64 armhf ppc64el riscv64 s390x"""
    sources = open("/etc/apt/sources.list").read()

    for version in Version:
        if version in sources:
            break

    print(f"version: {version}")

    content = f"""deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu-ports/ {version} main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu-ports/ {version} main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu-ports/ {version}-updates main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu-ports/ {version}-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu-ports/ {version}-backports main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu-ports/ {version}-backports main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu-ports/ {version}-security main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu-ports/ {version}-security main restricted universe multiverse"""

    conf = "/etc/apt/sources.list"
    writer(
        conf,
        content=content,
        read_only=False,
        official="https://mirrors.tuna.tsinghua.edu.cn/help/ubuntu-ports/",
    )
