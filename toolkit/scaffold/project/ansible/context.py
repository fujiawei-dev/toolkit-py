import uuid

from pydantic import BaseModel

from toolkit.logger import logging

log = logging.getLogger(__name__)


class AnsibleContext(BaseModel):
    install_location: str = "/usr/local/bin"
    executable_file: str = "main"
    config_file: str = "config.yaml"

    screen_id: str = None


ANSIBLE_USER_INPUT_CONTEXT = AnsibleContext().dict(exclude_none=True)


def ansible_user_input_context_hook(context: dict) -> dict:
    ansible_context = AnsibleContext(screen_id=uuid.uuid4().hex[4:16])

    ansible_user_input_context = AnsibleContext.parse_obj(context)

    return (
        context
        | ansible_context.dict(exclude_none=True)
        | ansible_user_input_context.dict(exclude_none=True)
    )
