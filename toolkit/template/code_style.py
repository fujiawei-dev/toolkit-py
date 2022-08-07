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
        text,  # 0 camel case
        text.replace(" ", "-"),  # 1 camel-case
        text.title(),  # 2 Camel Case
        text.replace(" ", "_"),  # 3 camel_case
    )
