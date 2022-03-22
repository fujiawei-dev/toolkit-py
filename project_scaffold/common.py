"""
Date: 2022.02.02 19:08
Description: Automatically create basic common files.
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 19:08
"""
import os
import shutil
from datetime import datetime
from pathlib import Path
from typing import List, Union

from unified_command.version import GENERATOR_HEADER
from .templates import TEMPLATES_COMMON_PATH

__all__ = [
    "create_common_files",
    "get_different_camel_case_styles",
    "Entity",
    "GENERATOR_HEADER",
]


def get_different_camel_case_styles(
    s: Union[str, Path] = Path.cwd()
) -> [str, str, str]:
    if isinstance(s, Path):
        s = s.stem
    else:
        s = str(s)

    s = list(s)
    for i in range(1, len(s)):
        if s[i - 1].islower() and s[i].isupper():
            s[i] = " " + s[i]

    package = "".join(s).lower().replace("_", "-").replace(" ", "-")  # camel-case
    package_title = package.replace("-", " ").title()  # Camel Case
    package_underscore = package.replace("-", "_")  # camel_case

    return package, package_title, package_underscore


class Entity(object):
    def __init__(self, file, content):
        self.file = file
        self.content = content.strip() + "\n"

    def create(self):
        if os.path.exists(self.file):
            return

        with open(self.file, "w", encoding="utf-8", newline="\n") as fp:
            fp.write(self.content)


# LICENSE
TEMPLATE_LICENSE = Entity(
    "LICENSE",
    (TEMPLATES_COMMON_PATH / "LICENSE")
    .read_text(encoding="utf-8")
    .format(CURRENT_YEAR=datetime.now().year),
)

# README.md
TEMPLATE_README = Entity("README.md", "# README\n")

# .gitignore
TEMPLATE_GITIGNORE = Entity(
    ".gitignore", (TEMPLATES_COMMON_PATH / ".gitignore").read_text(encoding="utf-8")
)


def create_common_files(folders: List[str] = None):
    TEMPLATE_LICENSE.create()
    TEMPLATE_README.create()
    TEMPLATE_GITIGNORE.create()

    if folders is None:
        folders = []

    folders.extend(
        [
            ".github/workflows",
            "assets",
            "examples",
            "drafts",
        ]
    )

    for folder in folders:
        common_folder = TEMPLATES_COMMON_PATH / folder
        if common_folder.exists():
            shutil.copytree(common_folder, folder, dirs_exist_ok=True)
        else:
            os.makedirs(folder, exist_ok=True)
