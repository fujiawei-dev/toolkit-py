"""The setup script."""

import os.path

from setuptools import find_packages, setup

import versioneer

with open("requirements.txt") as install_requires_file:
    install_requires = install_requires_file.read().splitlines()

with open("requirements-dev.txt") as dev_requires_file:
    dev_requires = dev_requires_file.read().splitlines()

with open("README.md") as readme_file:
    readme = readme_file.read()

with open("HISTORY.md") as history_file:
    history = history_file.read()


def find_package_data(*paths):
    return [
        os.path.normpath(os.path.join("..", p, f))
        for path in paths
        for (p, d, fs) in os.walk(path)
        for f in fs
    ]


setup(
    name="toolkit-py",
    version=versioneer.get_version(),
    cmdclass=versioneer.get_cmdclass(),
    author="Rustle Karl",
    author_email="fu.jiawei@outlook.com",
    url="https://github.com/fujiawei-dev/toolkit-py",
    description="Personal toolkit implemented by Python.",
    long_description=readme + "\n\n" + history,
    long_description_content_type="text/markdown",
    keywords="toolkit",
    license="MIT license",
    packages=find_packages(
        include=["toolkit", "toolkit.*"],
        exclude=(
            "tests",
            "tests.*",
            "docs",
            "toolkit.template.scaffold",
            "toolkit.template.scaffold.*",
        ),
    ),
    test_suite="tests",
    include_package_data=True,
    package_data={
        "toolkit": find_package_data("toolkit/template"),
    },
    entry_points={
        "console_scripts": [
            "toolkit=toolkit.cli:main",
            "mirror=toolkit.scaffold.mirror.cli:main",
            "config=toolkit.scaffold.config.cli:main",
            "project=toolkit.scaffold.project.cli:main",
            "gua=toolkit.cli:generate_user_agent_command",
            "upsfortypora=toolkit.cli:ups_for_typora_command",
        ],
    },
    python_requires=">=3.9",
    install_requires=install_requires,
    extras_require={"dev": dev_requires},
    classifiers=[
        "Environment :: Console",
        "Intended Audience :: Developers",
        "Operating System :: OS Independent",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Topic :: Software Development",
    ],
)
