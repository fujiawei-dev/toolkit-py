'''
Date: 2022.02.02 19:08
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 19:08
'''
from .common import Entity, TEMPLATE_MAKEFILE, create_common_files

PYTHON_MAKEFILE_CONTENT = TEMPLATE_MAKEFILE.content + '''
VERSION = 1.0.0
PACKAGE = package

all: test install clean

dep:
    pip install -r requirements.txt

build:
    python setup.py sdist
    python setup.py bdist_wheel

install: build
    pip install --force-reinstall --no-deps dist/$(PACKAGE)-$(VERSION).tar.gz

upload:
    twine upload dist/$(PACKAGE)-$(VERSION).tar.gz

test:
    pytest
    rm -r .pytest_cache

clean:
    rm -r build
    rm -r dist
    rm -r *egg-info
'''


def python():
    create_common_files(['tests', 'tests/data'])

    Entity('requirements.txt', '\n').create()
    Entity(TEMPLATE_MAKEFILE.file, PYTHON_MAKEFILE_CONTENT).create()
