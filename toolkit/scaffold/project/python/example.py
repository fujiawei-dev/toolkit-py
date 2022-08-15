from toolkit.config.context import USER_INPUT_CONTEXT
from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.template import TEMPLATE_PYTHON_PATH

create_example = generate_create_project_command(
    command_help="Create a python example project scaffold.",
    template_paths=TEMPLATE_PYTHON_PATH / "example",
    raw_user_input_context=USER_INPUT_CONTEXT,
)
