from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.template import TEMPLATE_DOCKER_PATH

create_compose = generate_create_project_command(
    command_help="Create a docker-compose project scaffold.",
    template_paths=TEMPLATE_DOCKER_PATH / "compose",
)
