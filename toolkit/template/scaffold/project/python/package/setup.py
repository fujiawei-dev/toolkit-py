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
    name="{{ project_slug.kebab_case }}",
    version=versioneer.get_version(),
    cmdclass=versioneer.get_cmdclass(),
    author="{{ author }}",
    author_email="{{ author_email }}",
    url="https://github.com/{{ github_username }}/{{ project_slug.kebab_case }}",
    description="{{ project_short_description }}",
    long_description=readme + "\n\n" + history,
    long_description_content_type="text/markdown",
    keywords="{{ project_slug.kebab_case }}",
    {% if open_source -%}
    license="MIT license",
    {%- endif %}
    packages=find_packages(
        include=["{{ project_slug.snake_case }}", "{{ project_slug.snake_case }}.*"],
        exclude=("tests", "tests.*", "docs"),
    ),
    test_suite="tests",
    include_package_data=True,
    entry_points={
        "console_scripts": [
            "{{ project_slug.snake_case }}={{ project_slug.snake_case }}.cli:main",
        ],
    },
    python_requires=">=3.9",
    install_requires=install_requires,
    extras_require={"dev": dev_requires},
    classifiers=[
        "Environment :: Console",
        "Intended Audience :: Developers",
        "Operating System :: OS Independent",
        {% if open_source -%}
        "License :: OSI Approved :: MIT License",
        {%- endif %}
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Topic :: Software Development",
    ],
)
