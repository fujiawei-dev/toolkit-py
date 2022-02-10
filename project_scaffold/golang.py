"""
Date: 2022.02.09 10:07
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.09 10:07
"""
from pathlib import Path

from .common import GENERATOR_HEADER, create_common_files, get_package
from .render import render_templates_recursively
from .templates import TEMPLATES_COMMON_PATH, TEMPLATES_PATH

GOLANG_HEADER = GENERATOR_HEADER.replace("#", "//")


def golang(include_suffixes=(".iris", ".cobra", ".viper")):
    package, package_title, package_underscore = get_package()

    templates = TEMPLATES_PATH / "golang"

    kwargs = {
        "PACKAGE_TITLE": package_title,
        "APP_NAME": package_underscore,
        "GOLANG_HEADER": GOLANG_HEADER,
        "GOLANG_MODULE": package_underscore,
        "MAKEFILE_HEADER": GENERATOR_HEADER,
    }

    render_templates_recursively(templates, Path.cwd(), include_suffixes, **kwargs)

    render_templates_recursively(
        TEMPLATES_COMMON_PATH, Path.cwd(), include_suffixes, **kwargs
    )

    create_common_files(["storage", "storage/configs"])
