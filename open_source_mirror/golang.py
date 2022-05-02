"""
Date: 2022.05.02 15:59
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.05.02 15:59
"""
from create_config_file.common import join_user, writer


def golang():
    content = """GO111MODULE=on
GOPROXY=https://goproxy.cn,direct
GOSUMDB=off
"""

    writer(
        join_user(".config/go/env"),
        content=content,
        read_only=False,
    )
