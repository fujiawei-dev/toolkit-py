"""
Date: 2022.02.03 12:37
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.03 12:37
"""
from .common import Entity
from .render import render_templates

TEMPLATE_ARTICLE_SETTINGS = Entity(
    "assets/templates/article_settings.yaml",
    """\
tags: ["标签"]
series: [ "系列"]
categories: [ "分类"]

toc: true  # 目录
draft: false  # 草稿
""",
)

TEMPLATE_ARTICLE_CONTENT = Entity(
    "assets/templates/article_content.md",
    """\
## 二级
### 三级
```shell
```
""".replace(
        "\n", "\n\n"
    ),
)


def notes():
    render_templates(
        "notes",
        folders=[
            "assets/images",  # 笔记配图
            "assets/templates",  # 笔记模板
            "docs",  # 基础语法
            "libraries",  # 库
            "libraries/standard",  # 标准库
            "libraries/tripartite",  # 第三方库
            "quickstart",  # 基础用法
            "refs",  # 参考中
            "src",  # 源码示例
            "src/docs",  # 基础语法源码示例
            "src/libraries/standard",  # 标准库源码示例
            "src/libraries/tripartite",  # 第三方库源码示例
            "src/quickstart",  # 基础用法源码示例
        ],
    )
