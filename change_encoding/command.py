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

TEXT_SUFFIXES = {
    ".asm",
    ".bat",
    ".c",
    ".cmd",
    ".cpp",
    ".go",
    ".h",
    ".ini",
    ".js",
    ".json",
    ".md",
    ".py",
    ".qml",
    ".qrc",
    ".ts",
    ".txt",
    ".ui",
    ".yaml",
    ".yml",
}

TEXT_FILES = {
    "license",
    "makefile",
}


def change_encoding(src: [Path, str], dst: Path = None, encoding="utf-8"):
    if not isinstance(src, Path):
        src = Path(src)

    if src.suffix.lower() in TEXT_SUFFIXES or src.stem.lower() in TEXT_FILES:
        dst = dst or src
        encoding = encoding or "utf-8"

        with open(src, "rb+") as fp:
            # CRLF -> LF
            content = fp.read().replace(b"\r\n", b"\n")
            original_encoding = chardet.detect(content)["encoding"]

            if not original_encoding:
                print(f"[error] {src} chardet.detect failed")
                return

            content = content.decode(original_encoding)
            content = (clean_characters(content).strip() + "\n").encode(encoding)

            if dst == src:
                fp.seek(0)
                fp.truncate()
                fp.write(content)

            else:
                with open(dst, "wb") as fw:
                    fw.write(content)

            print(f"{os.path.basename(src)}: {original_encoding} -> {encoding}")


def change_all_files_encoding(directory: Path = None, encoding="utf-8"):
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
