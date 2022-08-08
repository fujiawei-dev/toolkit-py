import os
from itertools import chain
from pathlib import Path
from typing import Any, Iterable, NamedTuple, Union

from pydantic import BaseModel

from toolkit.logger import logging

log = logging.getLogger(__name__)


class GolangCliFrameworks(NamedTuple):
    cli: str = ".cli"
    cobra: str = ".cobra"


class GolangLogFrameworks(NamedTuple):
    golog: str = ".golog"
    zerolog: str = ".zerolog"


class GolangWebFrameworks(NamedTuple):
    echo: str = ".echo"
    fiber: str = ".fiber"
    gin: str = ".gin"
    iris: str = ".iris"


class GolangConfigFrameworks(NamedTuple):
    viper: str = ".viper"


class GolangFrameworks(NamedTuple):
    nil: str = "nil"

    cli: Any = GolangCliFrameworks()
    log: Any = GolangLogFrameworks()
    web: Any = GolangWebFrameworks()
    config: Any = GolangConfigFrameworks()


F = GolangFrameworks()


class GolangContext(BaseModel):
    open_source: bool = True

    # cli_framework: Union[list, str] =F.cli
    # log_framework: Union[list, str] = F.log
    # web_framework: Union[list, str] = F.web

    frameworks: Union[list[list], list] = [
        [F.web.iris, F.log.golog, F.cli.cobra, F.config.viper],
        [F.web.fiber, F.log.zerolog, F.cli.cobra, F.config.viper],
        [F.web.echo, F.nil, F.cli.cobra, F.config.viper],
        [F.web.gin, F.log.zerolog, F.cli.cli, F.nil],
        [F.web.gin, F.log.golog, F.cli.cobra, F.config.viper],
    ]

    web_framework: str = None
    log_framework: str = None
    cli_framework: str = None
    config_framework: str = None

    log_error_f: str = None
    main_module: str = None

    web_framework_context: str = None
    web_framework_delete: str = None
    web_framework_engine: str = None
    web_framework_engine_group: str = None
    web_framework_error: str = None
    web_framework_get: str = None
    web_framework_import: str = None
    web_framework_jwt_down: str = None
    web_framework_jwt_up: str = None
    web_framework_nil: str = None
    web_framework_post: str = None
    web_framework_put: str = None
    web_framework_query_id: str = None
    web_framework_return: str = None
    web_framework_router_group: str = None


GOLANG_USER_INPUT_CONTEXT = GolangContext().dict(exclude_none=True)


def generate_wildcard_ignored_items(items: Iterable, ignored_items: list) -> list:
    wildcard_items = []

    for item in items:
        if item not in ignored_items:
            wildcard_items.append("*" + item + ".*")

    return wildcard_items


def golang_user_input_context_hook(context: dict) -> dict:
    # https://stackoverflow.com/questions/62267544/generate-pydantic-model-from-a-dict
    golang_context = GolangContext(
        log_error_f="Errorf",
        web_framework_delete="DELETE",
        web_framework_error="",
        web_framework_get="GET",
        web_framework_jwt_down="",
        web_framework_jwt_up="conf.JWTMiddleware(), ",
        web_framework_nil="",
        web_framework_post="POST",
        web_framework_put="PUT",
        web_framework_query_id=":id",
        web_framework_return="",
    )

    golang_user_input_context = GolangContext.parse_obj(context)

    golang_context.main_module = (
        f"https://github.com/{context['github_username']}/"
        if golang_user_input_context.open_source
        else ""
    ) + f"{context['project_slug']['kebab_case']}"

    (
        golang_context.web_framework,
        golang_context.log_framework,
        golang_context.cli_framework,
        golang_context.config_framework,
    ) = golang_user_input_context.frameworks

    context["cookiecutter"]["_ignore"].extend(
        generate_wildcard_ignored_items(
            chain(*(f for f in F if f != F.nil)),
            golang_user_input_context.frameworks,
        )
    )

    if golang_context.web_framework == F.web.echo:
        golang_context.web_framework_context = "echo.Context"
        golang_context.web_framework_error = "error"
        golang_context.web_framework_import = "github.com/labstack/echo/v4"
        golang_context.web_framework_jwt_down = ", conf.JWTMiddleware()"
        golang_context.web_framework_jwt_up = ""
        golang_context.web_framework_nil = "nil"
        golang_context.web_framework_return = "return"
        golang_context.web_framework_router_group = "echo.Group"
    elif golang_context.web_framework == F.web.fiber:
        golang_context.web_framework_context = "*fiber.Ctx"
        golang_context.web_framework_delete = "Delete"
        golang_context.web_framework_error = "error"
        golang_context.web_framework_get = "Get"
        golang_context.web_framework_import = "github.com/gofiber/fiber/v2"
        golang_context.web_framework_nil = "nil"
        golang_context.web_framework_post = "Post"
        golang_context.web_framework_put = "Put"
        golang_context.web_framework_return = "return"
        golang_context.web_framework_router_group = "fiber.Router"
    elif golang_context.web_framework == F.web.gin:
        golang_context.web_framework_context = "*gin.Context"
        golang_context.web_framework_engine = "*gin.Engine"
        golang_context.web_framework_engine_group = "Group"
        golang_context.web_framework_import = "github.com/gin-gonic/gin"
        golang_context.web_framework_router_group = "*gin.RouterGroup"
    elif golang_context.web_framework == F.web.iris:
        golang_context.web_framework_context = "iris.Context"
        golang_context.web_framework_delete = "Delete"
        golang_context.web_framework_engine = "*iris.Application"
        golang_context.web_framework_engine_group = "Party"
        golang_context.web_framework_get = "Get"
        golang_context.web_framework_import = "github.com/kataras/iris/v12"
        golang_context.web_framework_post = "Post"
        golang_context.web_framework_put = "Put"
        golang_context.web_framework_query_id = "{id:uint}"
        golang_context.web_framework_router_group = "iris.Party"

    if golang_context.log_framework == F.log.zerolog:
        golang_context.log_error_f = "Error().Msgf"

    log.debug(f"golang_context: {golang_context}")

    return (
        context
        | golang_context.dict(exclude_none=True)
        | golang_user_input_context.dict(exclude_none=True)
    )


def golang_generated_path_hook(path: Union[str, Path]) -> str:
    parent, filename = os.path.split(path)

    if (dot := filename.find(".")) != -1:
        stem, suffix = filename[:dot], filename[dot + 1 :]
        if (dot := suffix.rfind(".")) != -1:
            suffix = suffix[dot + 1 :]
    else:
        stem, suffix = filename, ""

    return os.path.join(parent, stem + "." + suffix)
