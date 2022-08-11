from pydantic import BaseModel

from toolkit.config.runtime import WINDOWS
from toolkit.logger import logging

log = logging.getLogger(__name__)

PYTHON_CONTEXT = {
    "plugins": [
        ("{{project_slug.snake_case}}/config/registry.py", WINDOWS),
    ],
}


class PythonContext(BaseModel):
    open_source: bool = True
    enable_click_group: bool = True
    enable_publish_action: bool = False


PYTHON_USER_INPUT_CONTEXT = PythonContext().dict(exclude_none=True)


def python_user_input_context_hook(context: dict) -> dict:
    python_context = PythonContext()

    python_user_input_context = PythonContext.parse_obj(context)

    ignored = []

    if not python_user_input_context.open_source:
        ignored.append("LICENSE")
    if not python_user_input_context.enable_publish_action:
        ignored.append(".github/workflows/python-publish.yml")

    context["cookiecutter"]["_ignore"].extend(ignored)

    return (
        context
        | python_context.dict(exclude_none=True)
        | python_user_input_context.dict(exclude_none=True)
    )
