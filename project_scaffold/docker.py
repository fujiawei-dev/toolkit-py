"""
Date: 2022.05.17 18:52
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.05.17 18:52
"""
from .render import render_templates


def docker():
    render_templates(
        "docker",
        common=False,
    )
