"""
Date: 2022.02.07 08:14
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.07 08:14
"""
import re

# 不可见字符
INVISIBLE_CHARACTERS = ("\xa0", "\x0b", "\x0c", "\u200b", "\u3000")

# 特殊空格字符
SPACE_CHARACTERS = ("&nbsp;",)

# 行末尾空格
PATTERN_LINE_RIGHT_WHITESPACE = re.compile(r"[ \t]+\n")


# 替换一组字符
def replace_all(string, words=tuple(), char="") -> str:
    """Replace a set of characters."""
    for w in words:
        string = string.replace(w, char)
    return string


def clean_characters(string: str) -> str:
    string = replace_all(string, INVISIBLE_CHARACTERS)
    string = replace_all(string, SPACE_CHARACTERS, " ")
    return PATTERN_LINE_RIGHT_WHITESPACE.sub("\n", string)
