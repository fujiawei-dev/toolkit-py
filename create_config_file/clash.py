"""
Date: 2022.02.03 12:20
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.03 12:20
"""
from .common import join_user, writer


def clash():
    writer(join_user(".config/clash/config.yaml"))
