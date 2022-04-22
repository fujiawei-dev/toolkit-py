"""
Date: 2022.02.09 10:07
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.09 10:07
"""
from enum import Enum

from .common import Entity, GENERATOR_HEADER, get_different_camel_case_styles
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


def golang(combination=GoCombinations.C2, entity=""):
    only_files = ""
    replace_list = {}

    if entity != "":
        only_files = "entity_template"

        (
            package,  # camel-case
            package_title,  # Camel Case
            package_underscore,  # camel_case
        ) = get_different_camel_case_styles(entity)

        replace_list = {
            "entity_template": package_underscore,
            "EntityTemplate": package_title.replace(" ", ""),
        }

    framework_values = GoCombinations.shortcuts(combination).split(";")

    kwargs = {
        "GET_STRING": "GET",
        "POST_STRING": "POST",
        "PUT_STRING": "PUT",
        "DELETE_STRING": "DELETE",
        "RETURN_STRING": "",
        "ERROR_STRING": "",
        "NIL_STRING": "",
        "QUERY_ID": ":id",
        "WEB_JWT_UP": "conf.JWTMiddleware(), ",
        "WEB_JWT_DOWN": "",
        "ERROR_F": "Errorf",
    }

    if GoWebFramework.Gin in framework_values:
        kwargs["WEB_FRAMEWORK"] = ".gin"
        kwargs["WEB_ENGINE"] = "*gin.Engine"
        kwargs["WEB_ENGINE_GROUP"] = "Group"
        kwargs["WEB_FRAMEWORK_IMPORT"] = "github.com/gin-gonic/gin"
        kwargs["ROUTER_GROUP"] = "*gin.RouterGroup"
        kwargs["WEB_CONTEXT"] = "*gin.Context"
    elif GoWebFramework.Echo in framework_values:
        kwargs["WEB_FRAMEWORK"] = ".echo"
        kwargs["WEB_FRAMEWORK_IMPORT"] = "github.com/labstack/echo/v4"
        kwargs["ROUTER_GROUP"] = "*echo.Group"
        kwargs["WEB_CONTEXT"] = "echo.Context"
        kwargs["RETURN_STRING"] = "return "
        kwargs["ERROR_STRING"] = "error"
        kwargs["NIL_STRING"] = " nil"
        kwargs["WEB_JWT_UP"] = ""
        kwargs["WEB_JWT_DOWN"] = ", conf.JWTMiddleware()"
    elif GoWebFramework.Fiber in framework_values:
        kwargs["WEB_FRAMEWORK"] = ".fiber"
        kwargs["WEB_FRAMEWORK_IMPORT"] = "github.com/gofiber/fiber/v2"
        kwargs["ROUTER_GROUP"] = "fiber.Router"
        kwargs["WEB_CONTEXT"] = "*fiber.Ctx"
        kwargs["ERROR_STRING"] = "error"
        kwargs["NIL_STRING"] = " nil"
        kwargs["RETURN_STRING"] = "return "
        kwargs["GET_STRING"] = "Get"
        kwargs["POST_STRING"] = "Post"
        kwargs["PUT_STRING"] = "Put"
        kwargs["DELETE_STRING"] = "Delete"
    elif GoWebFramework.Iris in framework_values:
        kwargs["WEB_FRAMEWORK"] = ".iris"
        kwargs["WEB_ENGINE"] = "*iris.Application"
        kwargs["WEB_ENGINE_GROUP"] = "Party"
        kwargs["WEB_FRAMEWORK_IMPORT"] = "github.com/kataras/iris/v12"
        kwargs["ROUTER_GROUP"] = "iris.Party"
        kwargs["WEB_CONTEXT"] = "iris.Context"
        kwargs["GET_STRING"] = "Get"
        kwargs["POST_STRING"] = "Post"
        kwargs["PUT_STRING"] = "Put"
        kwargs["DELETE_STRING"] = "Delete"
        kwargs["QUERY_ID"] = "{id:uint}"

    if GoCliFramework.Cli in framework_values:
        kwargs["CLI_FRAMEWORK"] = ".cli"

    elif GoCliFramework.Cobra in framework_values:
        kwargs["CLI_FRAMEWORK"] = ".cobra"

    if GoLogFramework.Zerolog in framework_values:
        kwargs["ERROR_F"] = "Error().Msgf"

    render_templates(
        "golang",
        include_suffixes=framework_values,
        folders=["storage", "storage/configs"],
        only_files=only_files,
        replace_list=replace_list,
        GOLANG_HEADER=GENERATOR_HEADER.replace("#", "//"),
        **kwargs
    )

    if entity == "":
        Entity(file="internal/.gitignore", content="example.go\n*_example.go").create()
