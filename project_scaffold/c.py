"""
Date: 2022.02.05 20:22
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.05 20:22
"""
from enum import Enum

from .render import render_templates


def c(only_files: str = ""):
    folders = [
        "cmake-build-debug",
        "cmake-build-release",
    ]

    if not only_files:
        folders.extend(
            [
                "include",
                "lib",
            ]
        )

    render_templates(
        "c",
        folders=folders,
        common=False,
        only_files=only_files,
    )


class Qt5Templates(str, Enum):
    Console = ".console"
    Gui = ".gui"
    Qml = ".qml"


class Qt5MakeType(str, Enum):
    QMake = ".qmake"
    CMake = ".cmake"


def qt5(template=Qt5Templates.Gui, only_files: str = ""):
    render_templates(
        "qt5",
        include_suffixes=[template],
        folders=[
            "assets",
            "assets/fonts",
            "assets/images",
            "cmake-build-debug",
            # "cmake-build-release",
            "include",
            "lib",
        ],
        common=False,
        template=template,
        only_files=only_files,
    )
