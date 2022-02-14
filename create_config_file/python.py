"""
Date: 2022.02.03 9:11
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.03 9:11
"""
import os
from enum import IntEnum

from .common import writer


def pypirc(read_only=True):
    official = "https://packaging.python.org/en/latest/specifications/pypirc/"

    conf = os.path.join(os.path.expanduser("~"), ".pypirc")

    content = """\
# https://pypi.org/manage/account/#API%20tokens

[distutils]
index-servers =
    pypi
    testpypi
    private-repository

[pypi]
username = __token__
password = <PyPI token>

[testpypi]
username = __token__
password = <TestPyPI token>

[private-repository]
repository = <private-repository URL>
username = <private-repository username>
password = <private-repository password>
"""

    writer(conf, content=content, read_only=read_only, official=official)


class Method(IntEnum):
    pypirc = 1

    @classmethod
    def func(cls, method):
        return {
            cls.pypirc: pypirc,
        }.get(method, pypirc)


def python(method=Method.pypirc, read_only=True):
    Method.func(method)(read_only)
