{{PYTHON_HEADER}}

import os.path

from setuptools import find_packages, setup

from {{PYTHON_MODULE}}.version import __version__

# What packages are required for this module to be executed?
requires = [
    "click",
    "click_aliases",
]

extras_require = {
    ':python_version < "3.7"': [
        "dataclasses",
    ],
}


def find_package_data(path):
    return [os.path.join("..", p, f) for (p, d, fs) in os.walk(path) for f in fs]


root = os.path.abspath(os.path.dirname(__file__))

# Import the README and use it as the long-description.
with open(os.path.join(root, "README.md"), encoding="utf-8") as f:
    long_description = f.read()

setup(
    name="{{PYPI_PACKAGE}}",
    version=__version__,
    url="https://github.com/fujiawei-dev/{{PYPI_PACKAGE}}",
    keywords=["package"],
    description="Personal package implemented by Python.",
    long_description=long_description,
    long_description_content_type="text/markdown",
    author="Rustle Karl",
    author_email="fujiawei@outlook.com",
    license="MIT",
    packages=find_packages(exclude=("tests", "tests.*")),
    install_requires=requires,
    extras_require=extras_require,
    package_data={"{{PYTHON_MODULE}}": find_package_data("{{PYTHON_MODULE}}/templates")},
    entry_points={
        "console_scripts": [
            "{{PYTHON_MODULE}}={{PYTHON_MODULE}}.command:main",
        ],
    },
    classifiers=[
        "Intended Audience :: Developers",
        "Environment :: Console",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: Implementation :: CPython",
        "Topic :: Software Development :: Libraries :: Python Modules",
    ],
)
