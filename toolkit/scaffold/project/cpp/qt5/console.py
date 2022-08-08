from toolkit.config.context import USER_INPUT_CONTEXT
from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.cpp.qt5.context import (
    CPP_QT5_USER_INPUT_CONTEXT,
    cpp_qt5_user_input_context_hook,
)
from toolkit.scaffold.project.template import TEMPLATE_CPP_QT5_PATH

create_console = generate_create_project_command(
    command_help="Create a cpp qt5 console project scaffold.",
    template_path=TEMPLATE_CPP_QT5_PATH / "console",
    raw_user_input_context=USER_INPUT_CONTEXT | CPP_QT5_USER_INPUT_CONTEXT,
    user_input_context_hook=cpp_qt5_user_input_context_hook,
)
