"""
Date: 2022.03.28 20:50
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.03.28 20:50
"""
from create_config_file.common import join_user, writer
from .templates import TEMPLATES_PATH


def maven():
    template = TEMPLATES_PATH / "java/maven/settings.xml"

    writer(
        join_user(".m2/settings.xml"),
        content=template.read_text(encoding="utf-8"),
        read_only=False,
    )


def java():
    maven()
