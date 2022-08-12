# {{ project_slug.words_capitalized }}

[![Code Style: black](https://img.shields.io/badge/codestyle-black-000000.svg)](https://github.com/psf/black)
[![Pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/{{ github_username }}/{{ project_slug.kebab_case }}/blob/main/.pre-commit-config.yaml)
{% if open_source -%}
[![License](https://img.shields.io/pypi/l/{{ project_slug.kebab_case }})](https://github.com/{{ github_username }}/{{ project_slug.kebab_case }}/blob/main/LICENSE)
{%- endif %}
{% if enable_publish_action -%}
[![Latest Version](https://img.shields.io/pypi/v/{{ project_slug.kebab_case }})](https://pypi.org/project/{{ project_slug.kebab_case }}/)
[![Supported Python versions](https://img.shields.io/pypi/pyversions/{{ project_slug.kebab_case }})](https://pypi.python.org/pypi/{{ project_slug.kebab_case }})
{%- endif %}
[![Python Test](https://github.com/{{ github_username }}/{{ project_slug.kebab_case }}/actions/workflows/python-test.yml/badge.svg)](https://github.com/{{ github_username }}/{{ project_slug.kebab_case }}/actions/workflows/python-test.yml)
{% if enable_publish_action -%}
[![Python Publish](https://github.com/{{ github_username }}/{{ project_slug.kebab_case }}/actions/workflows/python-publish.yml/badge.svg)](https://github.com/{{ github_username }}/{{ project_slug.kebab_case }}/actions/workflows/python-publish.yml)
{%- endif %}
> {{ project_short_description }}

## Installation
{% if enable_publish_action -%}
```shell
pip install -U {{ project_slug.kebab_case }}
```

If you are in China, you can use the following command to install the latest version:

```shell
pip install -U {{ project_slug.kebab_case }} -i https://pypi.douban.com/simple
```
{%- endif %}
Install it from source code:

```shell
pip install git+https://github.com/{{ github_username }}/{{ project_slug.kebab_case }}.git@main
```

## Usage

```shell
{{ project_slug.snake_case }} --help
```
