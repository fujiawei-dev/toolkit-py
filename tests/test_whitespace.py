"""
Date: 2022.02.07 09:44
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.07 09:44
"""
from change_encoding.whitespace import (
    clean_characters,
    INVISIBLE_CHARACTERS,
    SPACE_CHARACTERS,
)


def test_clean_whitespace():
    pairs = [
        (
            f"clean {''.join(INVISIBLE_CHARACTERS)} {''.join(SPACE_CHARACTERS)}\n",
            "clean\n",
        ),
    ]
    for pair in pairs:
        assert clean_characters(pair[0]) == pair[1]
