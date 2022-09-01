from pydantic import BaseModel

from toolkit.scaffold.project.command import generate_create_project_command
from toolkit.scaffold.project.golang.context import (
    GOLANG_USER_INPUT_CONTEXT,
    golang_generated_path_hook,
    golang_user_input_context_hook,
)
from toolkit.scaffold.project.template import TEMPLATE_GOLANG_PATH
from toolkit.template.code_style import CamelCaseStyle, get_camel_case_styles


class FactoryContext(BaseModel):
    entity: str = "entity"
    entity_slug: CamelCaseStyle = None

    entity_chinese: str = "对象"


def factory_user_input_context_hook(context: dict) -> dict:
    factory_user_input_context = FactoryContext.parse_obj(context)
    styles = get_camel_case_styles(factory_user_input_context.entity)
    factory_user_input_context.entity_slug = styles
    return factory_user_input_context.dict(exclude_none=True)


create_factory = generate_create_project_command(
    command_help="Generate files from templates for existing project",
    template_paths=TEMPLATE_GOLANG_PATH / "factory",
    generated_path_hook=golang_generated_path_hook,
    raw_user_input_context=GOLANG_USER_INPUT_CONTEXT,
    factory_user_input_context=FactoryContext().dict(exclude_none=True),
    factory_user_input_context_hook=factory_user_input_context_hook,
    user_input_context_hook=golang_user_input_context_hook,
    editors=["goland", "code"],
)
