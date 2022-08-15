from typing import Union

import click
from pydantic import BaseModel

from toolkit.config.context import USER_INPUT_CONTEXT
from toolkit.logger import logging
from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.golang.context import F
from toolkit.scaffold.project.template import TEMPLATE_GOLANG_PATH

log = logging.getLogger(__name__)

TEMPLATE_GOLANG_EXAMPLE_PATH = TEMPLATE_GOLANG_PATH / "example"


@click.group(help="Create a golang example project scaffold.")
def create_example():
    pass


create_example_console = generate_create_project_command(
    command_help="Create a golang console example project scaffold.",
    template_paths=TEMPLATE_GOLANG_EXAMPLE_PATH / "console",
)

create_example.add_command(create_example_console, "console")


class ExampleWebContext(BaseModel):
    web_framework: Union[str, list] = list(F.web)

    http_port: int = None


def example_web_user_input_context_hook(context: dict) -> dict:
    example_web_context = ExampleWebContext(
        http_port=26535,
    )

    example_web_user_input_context = ExampleWebContext.parse_obj(context)

    log.debug(f"example_web_context: {example_web_context}")

    return (
        context
        | example_web_context.dict(exclude_none=True)
        | example_web_user_input_context.dict(exclude_none=True)
    )


create_example_web = generate_create_project_command(
    command_help="Create a golang web example project scaffold.",
    template_paths=TEMPLATE_GOLANG_EXAMPLE_PATH / "web",
    raw_user_input_context=USER_INPUT_CONTEXT
    | ExampleWebContext().dict(exclude_none=True),
    user_input_context_hook=example_web_user_input_context_hook,
)


create_example.add_command(create_example_web, "web")
