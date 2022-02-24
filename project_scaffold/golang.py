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
    Echo = ".echo"


class LoggerFramework(str, Enum):
    Zerolog = ".zerolog"
    Golog = ".golog"


class CommandLineFramework(str, Enum):
    Cobra = ".cobra"


class ConfigurationManagementFramework(str, Enum):
    Viper = ".viper"


class Combinations(str, Enum):
    c1 = "1"
    c2 = "2"
    c3 = "3"

    C1 = ";".join(
        [
            HttpFramework.Iris,
            LoggerFramework.Golog,
            CommandLineFramework.Cobra,
            ConfigurationManagementFramework.Viper,
        ]
    )

    C2 = ";".join(
        [
            HttpFramework.Fiber,
            LoggerFramework.Zerolog,
            CommandLineFramework.Cobra,
            ConfigurationManagementFramework.Viper,
        ]
    )

    C3 = ";".join(
        [
            HttpFramework.Echo,
            CommandLineFramework.Cobra,
            ConfigurationManagementFramework.Viper,
        ]
    )

    @staticmethod
    def shortcuts(m: str) -> str:
        if m.isalnum():
            return {
                Combinations.c1: Combinations.C1,
                Combinations.c2: Combinations.C2,
                Combinations.c3: Combinations.C3,
            }.get(m, Combinations.C2)

        return m


def golang(combination=Combinations.C2):
    render_templates(
        "golang",
        include_suffixes=Combinations.shortcuts(combination).split(";"),
        folders=["storage", "storage/configs"],
        GOLANG_HEADER=GENERATOR_HEADER.replace("#", "//"),
    )
