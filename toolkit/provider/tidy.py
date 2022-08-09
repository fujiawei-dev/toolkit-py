import os
import re
from pathlib import Path
from typing import Callable, Iterable, List, Union

import chardet
import click
from binaryornot.check import is_binary

from toolkit.logger import logging

log = logging.getLogger(__name__)

# 不可见字符
INVISIBLE_CHARACTERS = ("\xa0", "\x0b", "\x0c", "\u200b", "\u3000")

# 特殊空格字符
SPACE_CHARACTERS = ("&nbsp;",)

# 行末尾空格
PATTERN_LINE_RIGHT_WHITESPACE = re.compile(r"[ \t]+\n")

# 常见二进制文件
BINARY_FILE_SUFFIXES = {".exe", ".dll", ".so", ".jpg", "jpeg", ".png", ".gif", ".pyc"}
BINARY_FILE_SUFFIXES |= {".7z", ".zip", ".lib"}


def is_binary_file(path: Union[str, Path]) -> bool:
    """Check if a file is binary."""
    # text_chars = bytearray({7, 8, 9, 10, 12, 13, 27} | set(range(0x20, 0x100)) - {0x7F})
    # return bool((open(path, "rb").read(1024)).translate(None, text_chars))
    return Path(path).suffix in BINARY_FILE_SUFFIXES or is_binary(str(path))


def replace_all(string: str, words: Iterable[str] = tuple(), char: str = "") -> str:
    """Replace a set of characters."""
    for w in words:
        string = string.replace(w, char)
    return string


def clean_characters(text: str) -> str:
    text = replace_all(text, INVISIBLE_CHARACTERS)
    text = replace_all(text, SPACE_CHARACTERS, " ")
    text = PATTERN_LINE_RIGHT_WHITESPACE.sub("\n", text)
    text = text.strip() + "\n"
    return text


def change_encoding(
    src: [Path, str],
    dst: [Path, str] = None,
    encoding: str = "utf-8",
    middlewares: list[Callable] = None,
):
    if dst and os.path.exists(dst):
        log.error(f"{dst} already exists")
        return

    if not is_binary_file(src):
        src, dst = Path(src), Path(dst or src)
        encoding = encoding or "utf-8"

        with open(src, "rb+") as fp:
            content = fp.read().replace(b"\r\n", b"\n")  # CRLF -> LF

            if not (original_encoding := chardet.detect(content)["encoding"]):
                original_encoding = "ascii"

            if original_encoding not in {"utf-8", "ascii"}:
                text = content.decode(original_encoding)

                if middlewares:
                    for middleware in middlewares:
                        text = middleware(text)

                content = text.encode(encoding)

                if dst == src:
                    fp.seek(0)
                    fp.truncate()
                    fp.write(content)
                else:
                    dst.write_bytes(content)

                log.info(f"{src.name}: {original_encoding} -> {encoding}")


def change_all_files_encoding(
    directory: Union[Path, str] = ".",
    encoding="utf-8",
    middlewares: List[Callable] = None,
):
    if not directory:
        directory = Path.cwd()
    elif not os.path.isdir(directory):
        return

    for item in Path(directory).iterdir():
        if item.is_dir():
            if item.stem[0] not in {".", "_"} or item.stem in {".github"}:
                change_all_files_encoding(item, encoding, middlewares)
        else:
            change_encoding(item, None, encoding, middlewares)


@click.command(help="Simple tidy/format the files, then change encoding.")
@click.argument("encoding", required=False, default="utf-8", type=click.STRING)
def tidy_command(encoding: str):
    change_all_files_encoding(encoding=encoding, middlewares=[clean_characters])
