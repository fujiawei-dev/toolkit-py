"""
Date: 2022.02.09 10:07
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.09 10:07
"""
from enum import Enum

from .common import GENERATOR_HEADER
from .render import render_templates


class HttpFramework(str, Enum):
    Iris = ".iris"
    Fiber = ".fiber"


class LoggerFramework(str, Enum):
    Zerolog = ".zerolog"


Combination1 = (HttpFramework.Iris, ".cobra", ".viper")
Combination2 = (HttpFramework.Fiber, LoggerFramework.Zerolog, ".cobra", ".viper")

Default = Combination2


def golang(include_suffixes=Default):
    render_templates(
        "golang",
        include_suffixes=include_suffixes,
        folders=["storage", "storage/configs"],
        GOLANG_HEADER=GENERATOR_HEADER.replace("#", "//"),
    )
