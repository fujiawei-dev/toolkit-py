from pydantic import BaseModel

from toolkit.logger import logging

log = logging.getLogger(__name__)

PYTHON_CONTEXT = {
    "plugins": [],
}


class PythonContext(BaseModel):
    open_source: bool = True
    enable_docker: bool = True
    enable_publish_action: bool = False
    enable_cli_command_group: bool = True


PYTHON_USER_INPUT_CONTEXT = PythonContext().dict(exclude_none=True)


def python_user_input_context_hook(context: dict) -> dict:
    python_context = PythonContext()

    python_user_input_context = PythonContext.parse_obj(context)

    ignored = []

    if not python_user_input_context.open_source:
        ignored.extend(["LICENSE", ".github/*", ".github"])
    if not python_user_input_context.enable_publish_action:
        ignored.append(".github/workflows/python-publish.yml")
    if not python_user_input_context.enable_docker:
        ignored.extend(["docker/*", "docker"])

    context["cookiecutter"]["_ignore"].extend(ignored)

    return (
        context
        | python_context.dict(exclude_none=True)
        | python_user_input_context.dict(exclude_none=True)
    )
