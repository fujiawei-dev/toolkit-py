from pydantic import BaseModel

from toolkit.config.context import USER_INPUT_CONTEXT
from toolkit.logger import logging
from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.template import TEMPLATE_CPP_QT5_PATH

log = logging.getLogger(__name__)


class QmlContext(BaseModel):
    qt_version: str = "5.12.6"
    qt_tool_chain: list = [
        ["mingw73_64", "mingw730_64"],
        ["mingw73_32", "mingw730_32"],
    ]

    x64_arch: bool = None
    qt_compile_version: str = None
    qt_tool_version: str = None


def qml_user_input_context_hook(context: dict) -> dict:
    qml_context = QmlContext(x64_arch=False)

    qml_user_input_context = QmlContext.parse_obj(context)

    (
        qml_context.qt_compile_version,
        qml_context.qt_tool_version,
    ) = qml_user_input_context.qt_tool_chain

    if "64" in qml_context.qt_compile_version:
        qml_context.x64_arch = True

    log.debug(f"context: {qml_context}")

    return (
        context
        | qml_context.dict(exclude_none=True)
        | qml_user_input_context.dict(exclude_none=True)
    )


create_qml = generate_create_project_command(
    command_help="Create a cpp qt5 qml project scaffold.",
    template_paths=TEMPLATE_CPP_QT5_PATH / "qml",
    raw_user_input_context=USER_INPUT_CONTEXT | QmlContext().dict(exclude_none=True),
    user_input_context_hook=qml_user_input_context_hook,
    editors=["clion"],
)
