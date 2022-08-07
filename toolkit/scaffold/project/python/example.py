from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.template import TEMPLATE_PYTHON_PATH

create_example = generate_create_project_command(
    command_help="Create a python example project scaffold.",
    template_path=TEMPLATE_PYTHON_PATH / "example",
)
