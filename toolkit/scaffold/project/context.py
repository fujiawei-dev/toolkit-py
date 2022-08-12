from cookiecutter.prompt import read_user_choice, read_user_variable, read_user_yes_no


def get_user_input_context(raw_user_input_context: dict = None) -> dict:
    raw_user_input_context = raw_user_input_context or {}

    for key, value in raw_user_input_context.items():
        # isinstance(True, int) is True
        if isinstance(value, bool):
            raw_user_input_context[key] = read_user_yes_no(
                'Are you sure "{}"?'.format(key),
                "y" if value else "n",
            )

        elif isinstance(value, (str, bytes, int, float)):
            raw_user_input_context[key] = type(value)(
                read_user_variable(
                    'Please enter a value for "{}"'.format(key),
                    value,
                )
            )

        elif isinstance(value, list):
            if len(value) == 1:
                raw_user_input_context[key] = value[0]
                continue

            raw_user_input_context[key] = read_user_choice(
                'Please select a value for "{}"'.format(key),
                value,
            )

    return raw_user_input_context


def get_ignored_items(project_context: dict = None, fields: list = None) -> list:
    project_context = project_context or {}
    fields = fields or []

    ignored_items = []

    for field in fields:
        for item, allowed in project_context[field]:
            if not allowed or read_user_yes_no(f"Ignore {item!r}?", "y"):
                ignored_items.append(item)

    return ignored_items
