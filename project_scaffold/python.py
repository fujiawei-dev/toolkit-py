'''
Date: 2022.02.02 19:08
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 19:08
'''
from pathlib import Path

from .common import Entity, TEMPLATE_MAKEFILE, TEMPLATE_README, create_common_files

PYTHON_MAKEFILE_CONTENT: str = TEMPLATE_MAKEFILE.content + '''
VERSION = 0.0.1
PACKAGE = {package}

all: install test clean

dep:
    pip install -r requirements.txt

build:
    python setup.py sdist
    python setup.py bdist_wheel

uninstall:
    pip uninstall -y $(PACKAGE)

install: uninstall build
    pip install --force-reinstall --no-deps dist/$(PACKAGE)-$(VERSION).tar.gz

upload: build
    twine upload dist/$(PACKAGE)-$(VERSION).tar.gz

test:
    pytest
    rm -r .pytest_cache

clean:
    rm -r build
    rm -r dist
    rm -r *egg-info

tag:
    git tag v$(VERSION)
    git push origin v$(VERSION)
'''.replace('    ', '\t')  # 4 whitespaces -> tab

PYTHON_README_CONTENT = '''\
# {title}

## Installation

```shell
pip install -U {package}
```

```shell
pip install -U {package} -i https://pypi.douban.com/simple
```

## Usage

```shell

```
'''


def python():
    Entity('requirements.txt', '\n').create()

    package = Path.cwd().stem

    Entity(
            TEMPLATE_README.file,
            PYTHON_README_CONTENT.format(
                    title=package.replace('-', ' ').title(),
                    package=package,
            ),
    ).create()

    Entity(
            TEMPLATE_MAKEFILE.file,
            PYTHON_MAKEFILE_CONTENT.format(package=package),
    ).create()

    create_common_files(['tests', 'tests/data'])
