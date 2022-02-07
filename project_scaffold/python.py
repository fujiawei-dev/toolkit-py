"""
Date: 2022.02.02 19:08
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.06 10:08:31
"""
from pathlib import Path

from unified_command.version import GENERATOR_HEADER
from .common import Entity, TEMPLATE_MAKEFILE, TEMPLATE_README, create_common_files

PYTHON_MAKEFILE_CONTENT: str = (
    TEMPLATE_MAKEFILE.content
    + """
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
""".replace(
        "    ", "\t"
    )
)  # 4 whitespaces -> tab

PYTHON_README_CONTENT = """\
# {title}

[![PyPI](https://img.shields.io/pypi/v/{package})](https://pypi.org/project/{package}/)

[![Python Test](https://github.com/fujiawei-dev/{package}/actions/workflows/python-test.yml/badge.svg)](https://github.com/fujiawei-dev/{package}/actions/workflows/python-test.yml)
[![Python Publish](https://github.com/fujiawei-dev/{package}/actions/workflows/python-publish.yml/badge.svg)](https://github.com/fujiawei-dev/{package}/actions/workflows/python-publish.yml)

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
"""

PYTHON_VERSION_CONTENT = (
    GENERATOR_HEADER
    + """
__version__ = '0.0.1'
"""
)

PYTHON_TEST = Entity(
    ".github/workflows/python-test.yml",
    GENERATOR_HEADER
    + """
# https://help.github.com/actions/language-and-framework-guides/using-python-with-github-actions

name: Python Test

on:
  push:
    branches: [ master, main ]
  pull_request:
    branches: [ master, main ]

jobs:
  build:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        python-version: ["3.6", "3.7", "3.8", "3.9", "3.10"]
    steps:
    - uses: actions/checkout@v2
    - name: Set up Python ${{ matrix.python-version }}
      uses: actions/setup-python@v2
      with:
        python-version: ${{ matrix.python-version }}
    - name: Install dependencies
      run: |
        python -m pip install --upgrade pip
        pip install flake8 pytest
        if [ -f requirements.txt ]; then pip install -r requirements.txt; fi
    - name: Lint with flake8
      run: |
        # stop the build if there are Python syntax errors or undefined names
        flake8 . --count --select=E9,F63,F7,F82 --show-source --statistics
        # exit-zero treats all errors as warnings. The GitHub editor is 127 chars wide
        flake8 . --count --exit-zero --max-complexity=10 --max-line-length=127 --statistics
    - name: Test with pytest
      run: |
        pytest
""",
)

PYTHON_TEST_CONF = Entity(
    "tests/confitest.py",
    GENERATOR_HEADER
    + """
def pytest_sessionfinish(session, exitstatus):
    if exitstatus == 5:
        session.exitstatus = 0
""",
)

PYTHON_PUBLISH = Entity(
    ".github/workflows/python-publish.yml",
    GENERATOR_HEADER
    + """
# https://help.github.com/en/actions/language-and-framework-guides/using-python-with-github-actions#publishing-to-package-registries

name: Python Publish

on:
  push:
    tags:
      - '*'

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Set up Python
      uses: actions/setup-python@v2
      with:
        python-version: '3.x'
    - name: Install dependencies
      run: |
        python -m pip install --upgrade pip
        pip install build
    - name: Build package
      run: python -m build
    - name: Publish package
      if: github.event_name == 'push' && startsWith(github.ref, 'refs/tags')
      uses: pypa/gh-action-pypi-publish@release/v1.5
      with:
        user: __token__
        password: ${{ secrets.PYPI_API_TOKEN }}
""",
)

PYTHON_SETUP_CONTENT = (
    GENERATOR_HEADER
    + """
import os

from setuptools import setup

from {underscores_package}.version import __version__

# What packages are required for this module to be executed?
requires = [
    'click',
]

# Import the README and use it as the long-description.
cwd = os.path.abspath(os.path.dirname(__file__))
with open(os.path.join(cwd, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()

setup(
        name='{package}',
        version=__version__,
        url='https://github.com/fujiawei-dev/{package}',
        packages=['package'],
        description='description',
        long_description=long_description,
        long_description_content_type='text/markdown',
        license='MIT',
        author='Rustle Karl',
        author_email='fu.jiawei@outlook.com',
        install_requires=requires,

        classifiers=[
            'Intended Audience :: Developers',
            'Environment :: Console',
            'License :: OSI Approved :: MIT License',
            'Operating System :: OS Independent',
            'Programming Language :: Python :: 3',
            'Programming Language :: Python :: Implementation :: CPython',
            'Topic :: Software Development :: Libraries :: Python Modules',
        ],
)
"""
)


def python():
    Entity("requirements.txt", "\n").create()

    package = Path.cwd().stem
    underscores_package = package.replace("-", "_")

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
            title=package.replace("-", " ").title(),
            package=package,
        ),
    ).create()

    create_common_files([underscores_package, "tests", "tests/data"])

    Entity(
        f"{underscores_package}/version.py",
        PYTHON_VERSION_CONTENT,
    ).create()

    PYTHON_TEST.create()
    PYTHON_TEST_CONF.create()

    Entity("tests/__init__.py", GENERATOR_HEADER).create()

    PYTHON_PUBLISH.create()

    Entity(
        "setup.py",
        PYTHON_SETUP_CONTENT.format(
            package=package,
            underscores_package=underscores_package,
        ),
    ).create()
