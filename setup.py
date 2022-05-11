"""
Date: 2020-09-21 23:48:26
LastEditors: Rustle Karl
Description: Distutils setup file for Toolkit-Py.
LastEditTime: 2022.05.10 11:29:42
"""

try:
    from setuptools import setup, find_packages
except:
    raise ImportError("setuptools is required to install toolkit-py!")

import os.path

from unified_command.version import __version__

root = os.path.abspath(os.path.dirname(__file__))

# Import the README and use it as the long-description.
def get_long_description():
    """Extract description from README.md, for PyPI's usage"""

    def process_ignore_tags(buffer):
        return "\n".join(
            x for x in buffer.split("\n")
            if "<!-- ignore_ppi -->" not in x
        )

    try:
        fpath = os.path.join(root, "README.md")
        with open(fpath, encoding="utf-8") as f:
            readme = f.read()
            desc = readme.partition("<!-- begin_ppi_description -->")[2]
            desc = desc.partition("<!-- end_ppi_description -->")[0]
            return process_ignore_tags(desc.strip())
    except IOError:
        return ""


# What packages are required for this module to be executed?
requires = [
    "chardet",
    "click",
    "click_aliases",
    "jinja2",
    "lxml",
    "pyyaml",
    "requests",
]

extras_require = {
    ':python_version < "3.7"': [
        "dataclasses",
    ],
}


def find_package_data(*paths):
    return [
        os.path.normpath(os.path.join("..", p, f))
        for path in paths
        for (p, d, fs) in os.walk(path)
        for f in fs
    ]


# https://packaging.python.org/guides/distributing-packages-using-setuptools/
setup(
    name="toolkit-py",
    version=__version__,
    url="https://github.com/fujiawei-dev/toolkit-py",
    keywords=["toolkit", "toolset"],
    description="Personal toolkit implemented by Python.",
    long_description=get_long_description(),
    long_description_content_type="text/markdown",
    author="Rustle Karl",
    author_email="fujiawei@outlook.com",
    license="BSD",
    packages=find_packages(exclude=("tests", "tests.*")),
    install_requires=requires,
    extras_require=extras_require,
    package_data={
        "create_config_file": find_package_data("create_config_file/templates"),
        "open_source_mirror": find_package_data("open_source_mirror/templates"),
        "project_scaffold": find_package_data("project_scaffold/templates"),
    },
    # Build starting scripts automatically
    entry_points={
        "console_scripts": [
            "gua=user_agent:command_gua",
            "cfm=open_source_mirror:command_cfm",
            "cps=project_scaffold:command_cps",
            "upsfortypora=image_hosting_service:command_ups_for_typora",
            "ccf=create_config_file:command_ccf",
            "cen=change_encoding:command_cen",
            "ucmd=unified_command.command:command_ucmd",
        ],
    },
    python_requires=">=3.7",
    classifiers=[
        "Environment :: Console",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: BSD License",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Topic :: Software Development",
    ],
)
