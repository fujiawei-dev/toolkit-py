"""
Date: 2022.05.09 12:44
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.05.09 12:44
"""
from .common import join_user, writer


"""
[http "https://github.com"]
	proxy = http://127.0.0.1:8118
"""


def clash():
    writer(join_user(".gitconfig"))
