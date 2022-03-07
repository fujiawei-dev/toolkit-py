"""
Date: 2022.02.05 20:22
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.05 20:22
"""
from enum import Enum

from .render import render_templates


def c():
    render_templates(
        "c",
        folders=[
            "include",
            "lib",
        ],
    )


class Qt5Templates(str, Enum):
    Console = ".console"
    Gui = ".gui"
    Qml = ".qml"


def qt5(template=Qt5Templates.Gui):
    render_templates(
        "qt5",
        include_suffixes=[template],
        folders=[
            "cmake-build-debug",
            "include",
            "lib",
        ],
        common=False,
    )
