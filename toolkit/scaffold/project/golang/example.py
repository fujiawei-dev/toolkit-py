from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.template import TEMPLATE_GOLANG_PATH

create_example = generate_create_project_command(
    command_help="Create a golang example project scaffold.",
    template_paths=TEMPLATE_GOLANG_PATH / "example",
)
