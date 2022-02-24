"""
Date: 2022.02.05 20:22
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.05 20:22
"""
from .render import render_templates


def c():
    render_templates(
        "c",
        folders=[
            "include",
            "lib",
        ],
    )


def qt5():
    render_templates(
        "qt5",
        folders=[
            "include",
            "lib",
        ],
    )
