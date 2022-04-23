"""
Date: 2022.02.06 08:39
Description: Version Information
LastEditors: Rustle Karl
LastEditTime: 2022.03.26 08:17:08
"""
from datetime import datetime

# https://packaging.python.org/en/latest/guides/single-sourcing-package-version/

__version__ = "0.17.6"

GENERATOR_HEADER = (
    f"# Generated by Toolkit-Py[v{__version__}] Generator. "
    f"Created at {datetime.now()}."
)
