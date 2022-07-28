"""
Date: 2022.07.27 13:55
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.07.27 13:55
"""
from pathlib import Path


def example(suffix: str):
    if not suffix.startswith("."):
        suffix = "." + suffix

    path = Path.cwd()

    (path / ("main" + suffix)).touch()
    (path / "README.md").touch()
