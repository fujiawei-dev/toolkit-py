from cookiecutter.prompt import read_user_variable, read_user_yes_no


def get_user_input_context(raw_context: dict = None) -> dict:
    raw_context = raw_context or {}

    for key, value in raw_context.items():
        raw_context[key] = read_user_variable(key, value)

    return raw_context


def get_ignored_items(context: dict = None, fields: list = None) -> list:
    context = context or {}
    fields = fields or []

    ignored_items = []

    for field in fields:
        for item, allowed in context[field]:
            if not allowed or read_user_yes_no(f"Ignore {item}?", "y"):
                ignored_items.append(item)

    return ignored_items
