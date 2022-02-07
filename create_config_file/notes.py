"""
Date: 2022.02.06 16:43
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.06 16:43
"""
import os
import time

from project_scaffold.notes import TEMPLATE_ARTICLE_CONTENT, TEMPLATE_ARTICLE_STATIC

_separator = "---\n"

_template_creator = """\
date: {date}
author: "Rustle Karl"
""".format(
    date=time.strftime("%Y-%m-%dT%H:%M:%S+08:00")
)

_template_article_dynamic = """\
title: "{title}"
url:  "{url}"  # 永久链接
"""


def new_article(path):
    if os.path.exists(path):
        return

    parent = os.path.dirname(path)
    if parent and not os.path.exists(parent):
        os.makedirs(parent, exist_ok=True)

    cwd = os.path.abspath(os.getcwd())
    prefix = os.path.basename(cwd)

    uri = os.path.splitext(path)[0]
    title = os.path.basename(uri)
    url = "/".join(["posts", prefix, uri])

    template_article = (
        _separator
        + _template_creator
        + "\n"
        + _template_article_dynamic.format(title=title, url=url)
    )

    if os.path.isfile(TEMPLATE_ARTICLE_STATIC.file):
        with open(TEMPLATE_ARTICLE_STATIC.file) as fp:
            content = fp.read()
            if content:
                template_article += content
            else:
                template_article += TEMPLATE_ARTICLE_STATIC.content

    template_article += _separator + "\n"

    if os.path.isfile(TEMPLATE_ARTICLE_CONTENT.file):
        with open(TEMPLATE_ARTICLE_CONTENT.file) as fp:
            content = fp.read()
            if content:
                template_article += (content + "\n\n") * 5
            else:
                template_article += (TEMPLATE_ARTICLE_CONTENT.content + "\n\n") * 5

    with open(path, "w", encoding="utf-8", newline="\n") as fp:
        fp.write(template_article)
