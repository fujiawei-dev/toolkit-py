from cookiecutter.prompt import read_user_variable, read_user_yes_no

from toolkit.config.context import USER_INPUT_CONTEXT
from toolkit.config.runtime import WINDOWS

PYTHON_CONTEXT = {
    "plugins": [
        ("config/registry.py", WINDOWS),
    ],
}


def get_user_input_context() -> dict:
    for key, value in USER_INPUT_CONTEXT.items():
        USER_INPUT_CONTEXT[key] = read_user_variable(key, value)

    return USER_INPUT_CONTEXT


def get_ignored_items() -> list:
    ignored_items = []

    for item, allowed in PYTHON_CONTEXT["plugins"]:
        if not allowed or read_user_yes_no(f"Ignore {item}?", "y"):
            ignored_items.append(item)

    return ignored_items
