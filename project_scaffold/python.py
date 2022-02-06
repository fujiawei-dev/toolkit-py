'''
Date: 2022.02.02 19:08
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.06 10:08:31
'''
from pathlib import Path

from unified_command.version import __header_
from .common import Entity, TEMPLATE_MAKEFILE, TEMPLATE_README, create_common_files

PYTHON_MAKEFILE_CONTENT: str = TEMPLATE_MAKEFILE.content + '''
VERSION := $(shell python -c "from {underscores_package}.version import __version__; print(__version__, end='')")
PACKAGE = {package}

all: reinstall test

version:
    echo $(VERSION)

dep:
    pip install -r requirements.txt

build:
    python setup.py sdist
    python setup.py bdist_wheel

uninstall:
    pip uninstall -y $(PACKAGE)

install: uninstall build
    pip install --force-reinstall --no-deps dist/$(PACKAGE)-$(VERSION).tar.gz

reinstall: install clean

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

PYTHON_VERSION_CONTENT = __header_ + '''
__version__ = '0.0.1'
'''


def python():
    Entity('requirements.txt', '\n').create()

    package = Path.cwd().stem
    underscores_package = package.replace('-', '_')

    Entity(
            TEMPLATE_MAKEFILE.file,
            PYTHON_MAKEFILE_CONTENT.format(
                    package=package,
                    underscores_package=underscores_package,
            ),
    ).create()

    Entity(
            TEMPLATE_README.file,
            PYTHON_README_CONTENT.format(
                    title=package.replace('-', ' ').title(),
                    package=package,
            ),
    ).create()

    create_common_files([underscores_package, 'tests', 'tests/data'])

    Entity(
            f'{underscores_package}/version.py',
            PYTHON_VERSION_CONTENT,
    ).create()
