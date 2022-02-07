"""
Date: 2022.02.04 10:55
Description: Recursively change the encoding of text files in the current folder.
LastEditors: Rustle Karl
LastEditTime: 2022.02.04 10:55
"""
import os
from pathlib import Path

import chardet
import click

from .whitespace import clean_characters

TEXT_EXTENSIONS = {
    ".txt",
    ".py",
    ".go",
    ".js",
    ".ts",
    ".md",
    ".ini",
    ".json",
    ".asm",
}


def change_encoding(src, dst=None, encoding="utf-8"):
    if os.path.splitext(src)[-1] in TEXT_EXTENSIONS:
        if not dst:
            dst = src

        if not encoding:
            encoding = "utf-8"

        with open(src, "rb+") as fp:
            # CRLF -> LF
            content = fp.read().replace(b"\r\n", b"\n")
            original_encoding = chardet.detect(content)["encoding"]
            content = content.decode(original_encoding)
            content = clean_characters(content).encode(encoding)

            if dst == src:
                fp.seek(0)
                fp.truncate()
                fp.write(content)

            else:
                with open(dst, "wb") as fw:
                    fw.write(content)

            print(f"{os.path.basename(src)}: {original_encoding} -> {encoding}")


def change_all_files_encoding(directory=None, encoding="utf-8"):
    if directory is None:
        directory = Path.cwd()
    elif not os.path.isdir(directory):
        return

    for file in directory.iterdir():
        if file.is_dir():
            change_all_files_encoding(file, encoding=encoding)
        else:
            change_encoding(file, encoding=encoding)


@click.command()
@click.argument("encoding", required=False)
def command_cen(encoding):
    """All files in the current path are converted to the specified encoding."""
    change_all_files_encoding(encoding=encoding)
