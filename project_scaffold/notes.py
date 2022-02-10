"""
Date: 2022.02.03 12:37
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.03 12:37
"""
from .render import render_templates


def notes():
    render_templates(
        "notes",
        folder=[
            "quickstart",  # 安装/卸载/基础用法
            "docs",  # 基础语法/知识/原理
            "libraries",  # 库
            "libraries/standard",  # 标准库
            "libraries/tripartite",  # 第三方库
            "assets/images",  # 图片
            "assets/templates",  # 笔记模板
            "src",  # 源码示例
            "src/docs",  # 源码示例
            "src/quickstart",  # 源码示例
            "src/libraries/standard",  # 标准库源码示例
            "src/libraries/tripartite",  # 第三方库源码示例
        ],
    )
