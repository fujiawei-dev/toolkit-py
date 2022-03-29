"""
Date: 2022.03.29 12:58
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.03.29 12:58
"""
from project_scaffold.templates import TEMPLATES_PATH

from project_scaffold.render import render_by_jinja2


def test_render_by_jinja2():
    template = TEMPLATES_PATH / "golang/internal/api/api.go"

    print(
        render_by_jinja2(template.read_text(encoding="utf-8"), web_framework=".echo"),
        file=open("api.go", "w", encoding="utf-8"),
    )
