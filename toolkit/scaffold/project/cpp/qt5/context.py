from typing import Union

from pydantic import BaseModel

from toolkit.logger import logging

log = logging.getLogger(__name__)


class CppQt5Context(BaseModel):
    open_source: bool = True

    cpp_standard: Union[list, int] = [17, 20]

    qt_version: str = "5.12"
    qt_tool_chain: list = [
        ["mingw73_64", "mingw730_64"],
        ["mingw73_32", "mingw730_32"],
    ]

    enable_network: bool = False
    enable_event_loop: bool = False

    x64_arch: bool = None

    qt_compile_version: str = None
    qt_tool_version: str = None


CPP_QT5_USER_INPUT_CONTEXT = CppQt5Context().dict(exclude_none=True)


def cpp_qt5_user_input_context_hook(context: dict) -> dict:
    cpp_qt5_context = CppQt5Context(x64_arch=False)

    cpp_qt5_user_input_context = CppQt5Context.parse_obj(context)

    (
        cpp_qt5_context.qt_compile_version,
        cpp_qt5_context.qt_tool_version,
    ) = cpp_qt5_user_input_context.qt_tool_chain

    if "64" in cpp_qt5_context.qt_compile_version:
        cpp_qt5_context.x64_arch = True

    log.debug(f"cpp_qt5_context: {cpp_qt5_context}")

    return (
        context
        | cpp_qt5_context.dict(exclude_none=True)
        | cpp_qt5_user_input_context.dict(exclude_none=True)
    )
