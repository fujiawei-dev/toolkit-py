"""
Date: 2022.02.09 10:07
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.09 10:07
"""
from enum import Enum

from .common import Entity, GENERATOR_HEADER
from .render import render_templates


class GoWebFramework(str, Enum):
    Echo = ".echo"
    Fiber = ".fiber"
    Gin = ".gin"
    Iris = ".iris"


class GoLogFramework(str, Enum):
    Golog = ".golog"
    Zerolog = ".zerolog"


class GoCliFramework(str, Enum):
    Cli = ".cli"
    Cobra = ".cobra"


class GoConfigFramework(str, Enum):
    Viper = ".viper"


class GoCombinations(str, Enum):
    c1 = "1"
    c2 = "2"
    c3 = "3"
    c4 = "4"
    c5 = "5"

    C1 = ";".join(
        [
            GoWebFramework.Iris,
            GoLogFramework.Golog,
            GoCliFramework.Cobra,
            GoConfigFramework.Viper,
        ]
    )

    C2 = ";".join(
        [
            GoWebFramework.Fiber,
            GoLogFramework.Zerolog,
            GoCliFramework.Cobra,
            GoConfigFramework.Viper,
        ]
    )

    C3 = ";".join(
        [
            GoWebFramework.Echo,
            GoCliFramework.Cobra,
            GoConfigFramework.Viper,
        ]
    )

    C4 = ";".join(
        [
            GoWebFramework.Gin,
            GoLogFramework.Zerolog,
            GoCliFramework.Cli,
        ]
    )

    C5 = ";".join(
        [
            GoWebFramework.Gin,
            GoLogFramework.Zerolog,
            GoCliFramework.Cobra,
            GoConfigFramework.Viper,
        ]
    )

    @staticmethod
    def shortcuts(m: str) -> str:
        if not m:
            return GoCombinations.C1

        if m.isalnum():
            return {
                GoCombinations.c1: GoCombinations.C1,
                GoCombinations.c2: GoCombinations.C2,
                GoCombinations.c3: GoCombinations.C3,
                GoCombinations.c4: GoCombinations.C4,
                GoCombinations.c5: GoCombinations.C5,
            }.get(m, GoCombinations.C2)

        return m


def golang(combination=GoCombinations.C2):
    render_templates(
        "golang",
        include_suffixes=GoCombinations.shortcuts(combination).split(";"),
        folders=["storage", "storage/configs"],
        GOLANG_HEADER=GENERATOR_HEADER.replace("#", "//"),
    )

    Entity(file="internal/.gitignore", content="example.go\n*_example.go").create()
