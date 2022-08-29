from typing import NamedTuple


class CamelCaseStyle(NamedTuple):
    words_lowercase: str  # camel case
    kebab_case: str  # camel-case
    words_capitalized: str  # Camel Case
    snake_case: str  # camel_case
    pascal_case: str  # CamelCase


def get_camel_case_styles(text: str) -> CamelCaseStyle:
    symbols = "_- "

    chars = list(text.strip(symbols))

    for i in range(1, len(chars)):
        if chars[i] in symbols:
            chars[i] = " "
        elif chars[i - 1].islower() and chars[i].isupper():
            chars[i] = " " + chars[i]

    text = "".join(chars).lower()  # camel case

    return CamelCaseStyle(
        words_lowercase=text,  # camel case
        kebab_case=text.replace(" ", "-"),  # camel-case
        words_capitalized=text.title(),  # Camel Case
        snake_case=text.replace(" ", "_"),  # camel_case
        pascal_case=text.title().replace(" ", ""),  # CamelCase
    )
