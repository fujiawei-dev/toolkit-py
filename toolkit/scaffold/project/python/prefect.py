from toolkit.config.context import USER_INPUT_CONTEXT
from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.python.context import (
    PYTHON_CONTEXT,
    PYTHON_USER_INPUT_CONTEXT,
    python_user_input_context_hook,
)
from toolkit.scaffold.project.template import TEMPLATE_PYTHON_PATH

create_prefect = generate_create_project_command(
    command_help="Create a python prefect project scaffold.",
    template_paths=[TEMPLATE_PYTHON_PATH / "package", TEMPLATE_PYTHON_PATH / "prefect"],
    raw_user_input_context=USER_INPUT_CONTEXT | PYTHON_USER_INPUT_CONTEXT,
    user_input_context_hook=python_user_input_context_hook,
    project_context=PYTHON_CONTEXT,
    ignored_fields=["plugins"],
)
