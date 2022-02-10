"""
Date: 2022.02.09 10:07
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.09 10:07
"""
from .common import GENERATOR_HEADER
from .render import render_templates


def golang(include_suffixes=(".iris", ".cobra", ".viper")):
    render_templates(
        "golang",
        include_suffixes=include_suffixes,
        folder=["storage", "storage/configs"],
        GOLANG_HEADER=GENERATOR_HEADER.replace("#", "//"),
    )
