"""
Date: 2022.03.14 00:23
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.03.14 00:23
"""
from pathlib import Path


def remove_youdao_note_ad():
    """移除有道云笔记的广告，必须以管理员权限运行"""
    p = Path("C:/Program Files (x86)/Youdao/YoudaoNote/theme/build.xml")
    p.write_text(p.read_text().replace("161", "0").replace("250,160", "0,0"))
