def get_camel_case_styles(text: str) -> tuple[str, str, str, str]:
    symbols = "_- "

    chars = list(text.strip(symbols))

    for i in range(1, len(chars)):
        if chars[i] in symbols:
            chars[i] = " "
        elif chars[i - 1].islower() and chars[i].isupper():
            chars[i] = " " + chars[i]

    text = "".join(chars).lower()  # camel case

    return (
        text,  # camel case
        text.replace(" ", "-"),  # camel-case
        text.title(),  # Camel Case
        text.replace(" ", "_"),  # camel_case
    )
