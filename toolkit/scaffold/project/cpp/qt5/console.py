from typing import Union

from pydantic import BaseModel

from toolkit.config.context import USER_INPUT_CONTEXT
from toolkit.logger import logging
from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.template import TEMPLATE_CPP_QT5_PATH

log = logging.getLogger(__name__)


class ConsoleContext(BaseModel):
    cpp_standard: Union[list, int] = [17, 20]

    qt_version: str = "5.12.6"
    qt_tool_chain: list = [
        ["mingw73_64", "mingw730_64"],
        ["mingw73_32", "mingw730_32"],
    ]

    all_in_main: bool = True
    enable_network: bool = True
    enable_http_request: bool = False
    enable_websocket: bool = False
    enable_event_loop: bool = False
    enable_src_module: bool = False
    enable_3rd_module: bool = False

    x64_arch: bool = None
    qt_compile_version: str = None
    qt_tool_version: str = None


CONSOLE_USER_INPUT_CONTEXT = ConsoleContext().dict(exclude_none=True)


def console_user_input_context_hook(context: dict) -> dict:
    console_context = ConsoleContext(x64_arch=False)

    console_user_input_context = ConsoleContext.parse_obj(context)

    (
        console_context.qt_compile_version,
        console_context.qt_tool_version,
    ) = console_user_input_context.qt_tool_chain

    if "64" in console_context.qt_compile_version:
        console_context.x64_arch = True

    console_context.enable_network = any(
        [
            console_context.enable_http_request,
            console_context.enable_websocket,
        ]
    )

    console_context.enable_event_loop = any(
        [
            console_context.enable_event_loop,
            console_context.enable_network,
            console_context.enable_websocket,
        ]
    )

    ignored = []

    if not console_context.enable_src_module:
        ignored.append("src")

    context["cookiecutter"]["_ignore"].extend(ignored)

    log.debug(f"context: {console_context}")

    return (
        context
        | console_context.dict(exclude_none=True)
        | console_user_input_context.dict(exclude_none=True)
    )


create_console = generate_create_project_command(
    command_help="Create a cpp qt5 console project scaffold.",
    template_paths=TEMPLATE_CPP_QT5_PATH / "console",
    raw_user_input_context=USER_INPUT_CONTEXT | CONSOLE_USER_INPUT_CONTEXT,
    user_input_context_hook=console_user_input_context_hook,
)
