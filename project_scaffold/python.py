"""
Date: 2022.02.02 19:08
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.06 10:08:31
"""
from .common import GENERATOR_HEADER
from .render import render_templates


def python():
    render_templates(
        "python",
        special_paths=[".github"],
        PYTHON_HEADER=GENERATOR_HEADER,
    )
