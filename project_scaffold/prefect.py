"""
Date: 2022.07.01 10:33
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.07.01 10:33
"""
from project_scaffold.render import render_templates


def prefect():
    render_templates(
        "prefect",
        common=False,
    )
