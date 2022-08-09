from toolkit.config.context import USER_INPUT_CONTEXT
from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.golang.context import (
    GOLANG_USER_INPUT_CONTEXT,
    golang_generated_path_hook,
    golang_user_input_context_hook,
)
from toolkit.scaffold.project.template import TEMPLATE_GOLANG_PATH

create_web = generate_create_project_command(
    command_help="Create a golang web project scaffold.",
    template_paths=TEMPLATE_GOLANG_PATH / "web",
    generated_path_hook=golang_generated_path_hook,
    raw_user_input_context=USER_INPUT_CONTEXT | GOLANG_USER_INPUT_CONTEXT,
    user_input_context_hook=golang_user_input_context_hook,
)
