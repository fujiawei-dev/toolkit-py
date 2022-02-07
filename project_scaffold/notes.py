"""
Date: 2022.02.03 12:37
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.03 12:37
"""
from .common import Entity, create_common_files

TEMPLATE_ARTICLE_STATIC = Entity(
    "assets/templates/article_static.yaml",
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
    create_common_files(
        [
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
        ]
    )

    TEMPLATE_ARTICLE_STATIC.create()
    TEMPLATE_ARTICLE_CONTENT.create()
