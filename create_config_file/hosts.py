"""
Date: 2022.02.03 12:21
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.03 12:21
"""
from .common import is_windows, writer


def hosts():
    if is_windows():
        conf = "C:\\Windows\\System32\\drivers\\etc\\hosts"
    else:
        conf = "/etc/hosts"

    writer(conf)
