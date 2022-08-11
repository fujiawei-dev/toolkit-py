from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.golang.context import (
    GOLANG_USER_INPUT_CONTEXT,
    golang_generated_path_hook,
    golang_user_input_context_hook,
)
from toolkit.scaffold.project.template import TEMPLATE_GOLANG_PATH

create_factory = generate_create_project_command(
    command_help="Generate files from templates for existing project",
    template_paths=TEMPLATE_GOLANG_PATH / "factory",
    generated_path_hook=golang_generated_path_hook,
    raw_user_input_context=GOLANG_USER_INPUT_CONTEXT,
    factory_user_input_context={
        "entity": "entity",
        "entity_chinese": "对象",
    },
    user_input_context_hook=golang_user_input_context_hook,
)
