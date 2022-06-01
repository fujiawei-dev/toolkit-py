"""
Date: 2022.06.01 14:42
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.06.01 14:42
"""
from .render import render_templates


def qml():
    render_templates(
        "qt5",
        folders=[
            "assets",
            "assets/fonts",
            "assets/images",
            "frontend",
            "include",
            "include",
            "lib",
            "tests",
        ],
        common=False,
    )
