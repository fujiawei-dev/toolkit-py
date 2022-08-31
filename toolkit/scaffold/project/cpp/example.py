from pydantic import BaseModel

from toolkit.config.context import USER_INPUT_CONTEXT
from toolkit.logger import logging
from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.template import TEMPLATE_CPP_PATH

log = logging.getLogger(__name__)


class ExampleContext(BaseModel):
    x64_arch: bool = True
    enable_3rd_module: bool = False


create_example = generate_create_project_command(
    command_help="Create a cpp qt5 qml example project scaffold.",
    template_paths=TEMPLATE_CPP_PATH / "example",
    raw_user_input_context=USER_INPUT_CONTEXT
    | ExampleContext().dict(exclude_none=True),
    editors=["clion", "code"],
)
