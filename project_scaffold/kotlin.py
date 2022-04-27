"""
Date: 2022.04.27 22:18
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.04.27 22:18
"""
from .render import render_templates


def kotlin():
    render_templates(
        "kotlin",
        common=False,
    )
