import click
from pydantic import BaseModel

from toolkit.config.context import USER_INPUT_CONTEXT
from toolkit.logger import logging
from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.template import TEMPLATE_CPP_QT5_PATH

log = logging.getLogger(__name__)

TEMPLATE_CPP_QT5_EXAMPLE_PATH = TEMPLATE_CPP_QT5_PATH / "example"


@click.group(help="Create a cpp qt5 example project scaffold.")
def create_example():
    pass


class ExampleContext(BaseModel):
    qt_version: str = "5.12.6"
    qt_tool_chain: list = [
        ["mingw73_64", "mingw730_64"],
        ["mingw73_32", "mingw730_32"],
    ]

    enable_3rd_module: bool = False

    x64_arch: bool = None
    qt_compile_version: str = None
    qt_tool_version: str = None


def example_user_input_context_hook(context: dict) -> dict:
    example_context = ExampleContext(x64_arch=False)

    example_user_input_context = ExampleContext.parse_obj(context)

    (
        example_context.qt_compile_version,
        example_context.qt_tool_version,
    ) = example_user_input_context.qt_tool_chain

    if "64" in example_context.qt_compile_version:
        example_context.x64_arch = True

    log.debug(f"example_context: {example_context}")

    return (
        context
        | example_context.dict(exclude_none=True)
        | example_user_input_context.dict(exclude_none=True)
    )


create_example_qml = generate_create_project_command(
    command_help="Create a cpp qt5 qml example project scaffold.",
    template_paths=TEMPLATE_CPP_QT5_EXAMPLE_PATH / "qml",
    raw_user_input_context=USER_INPUT_CONTEXT
    | ExampleContext().dict(exclude_none=True),
    user_input_context_hook=example_user_input_context_hook,
)

create_example.add_command(create_example_qml, "qml")

create_example_console = generate_create_project_command(
    command_help="Create a cpp qt5 console example project scaffold.",
    template_paths=TEMPLATE_CPP_QT5_EXAMPLE_PATH / "console",
    raw_user_input_context=USER_INPUT_CONTEXT
    | ExampleContext().dict(exclude_none=True),
    user_input_context_hook=example_user_input_context_hook,
)

create_example.add_command(create_example_console, "console")
