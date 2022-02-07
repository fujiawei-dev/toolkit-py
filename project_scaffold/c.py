"""
Date: 2022.02.05 20:22
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.05 20:22
"""
from pathlib import Path

from .common import Entity, TEMPLATE_MAKEFILE, TEMPLATE_README, create_common_files

C_MAKEFILE_CONTENT: str = (
    TEMPLATE_MAKEFILE.content
    + """
PROJECT = {project}

TARGET = $(PROJECT)

INSTALL_PATH = /usr/local/bin/

ifeq ($(OS), Windows_NT)
TARGET = $(PROJECT).exe
INSTALL_PATH = c:/developer/bin/
endif

all: reinstall

build:
    gcc -o $(TARGET) $(PROJECT).c

uninstall:
    -(rm $(INSTALL_PATH)$(TARGET))

install:
    copy $(TARGET) $(INSTALL_PATH)
    rm $(TARGET)

reinstall: uninstall install

tag:
    git tag v$(VERSION)
    git push origin v$(VERSION)
""".replace(
        "    ", "\t"
    )
)  # 4 whitespaces -> tab

C_README_CONTENT = """\
# {title}

## Installation

```shell
make install
```

## Usage

```shell

```
"""


def c():
    project = Path.cwd().stem

    Entity(
        TEMPLATE_MAKEFILE.file,
        C_MAKEFILE_CONTENT.format(project=project.replace("-", "_")),
    ).create()

    Entity(
        TEMPLATE_README.file,
        C_README_CONTENT.format(
            title=project.replace("-", " ").title(),
            project=project,
        ),
    ).create()

    create_common_files()
