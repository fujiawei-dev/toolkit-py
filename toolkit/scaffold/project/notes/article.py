import os.path
import time
from pathlib import Path
from typing import NamedTuple, Union

import click
from cookiecutter.prompt import read_user_choice, read_user_yes_no

from toolkit.scaffold.project import cpp, golang, python
from toolkit.template.code_style import get_camel_case_styles

ARTICLE_SETTINGS_PATH = "assets/templates/article_settings.yaml"

ARTICLE_CONTENT_PATH = "assets/templates/article_content.md"

ARTICLE_HEADER_SEPARATOR = "---\n"

ARTICLE_HEADER_CREATOR = f"""\
date: {time.strftime("%Y-%m-%dT%H:%M:%S+08:00")}  # 创建日期
author: "Rustle Karl"  # 作者
"""

ARTICLE_HEADER_TITLE = """\
title: "{title}"  # 文章标题
url:  "{url}"  # 设置网页永久链接
"""


class ProgrammingLanguage(NamedTuple):
    none: str = "none"

    python: str = "python"
    golang: str = "golang"
    cpp: str = "cpp"
    kotlin: str = "kotlin"


L = ProgrammingLanguage()


def render_article_content(
    article_path: Union[str, Path],
    workspace_path: str = os.getcwd(),
    header_only: bool = False,
) -> str:
    prefix = os.path.basename(workspace_path)

    uri = os.path.splitext(article_path)[0]
    title = os.path.basename(uri)
    url = os.path.normpath("/".join(["posts", prefix, uri])).replace("\\", "/")

    project_slugs = get_camel_case_styles(title)

    article_content = (
        ARTICLE_HEADER_SEPARATOR
        + ARTICLE_HEADER_CREATOR
        + "\n"
        + ARTICLE_HEADER_TITLE.format(title=project_slugs[2], url=url)
    )

    if os.path.isfile(ARTICLE_SETTINGS_PATH):
        with open(ARTICLE_SETTINGS_PATH, encoding="utf-8") as fp:
            article_content += fp.read().format(tag=project_slugs[1])

    article_content += ARTICLE_HEADER_SEPARATOR + "\n"

    if not header_only and os.path.isfile(ARTICLE_CONTENT_PATH):
        with open(ARTICLE_CONTENT_PATH, encoding="utf-8") as fp:
            article_content += (fp.read() + "\n\n") * 3

    return article_content.strip() + "\n"


@click.command(help="Create a new article.")
@click.option(
    "--article-path",
    "-p",
    type=click.STRING,
    required=True,
    help="The path of the article or articles folder.",
)
@click.option(
    "--workspace-path",
    type=click.STRING,
    default=os.getcwd(),
    help="The path of the workspace.",
)
@click.option("--header-only", is_flag=True, help="Only create the header.")
@click.pass_context
def create_article(
    ctx: click.Context,
    article_path: str,
    workspace_path: str,
    header_only: bool,
):
    if not os.path.exists(article_path) or header_only:
        if header_only:
            article_paths = []
            article_path = Path(article_path)

            if article_path.is_file() and article_path.suffix == ".md":
                article_paths.append(article_path)
            elif article_path.is_dir():
                article_paths.extend(article_path.glob("*.md"))

            for article_path in article_paths:
                if article_path.open().read(3) == "---":
                    continue

                old_content = article_path.read_text(encoding="utf-8")

                with article_path.open(mode="w", encoding="utf-8", newline="\n") as fp:
                    fp.write(
                        render_article_content(article_path, workspace_path, True)
                        + old_content
                    )

        else:
            parent, filename = os.path.split(article_path)
            os.makedirs(parent, exist_ok=True)

            if "." not in filename:
                article_path += ".md"

            with open(article_path, "w", encoding="utf-8") as fp:
                fp.write(render_article_content(article_path, workspace_path))

        language = read_user_choice("Language", list(L))

        src_project_path = os.path.join("src", os.path.splitext(article_path)[0])

        if language == L.python:
            ctx.invoke(
                python.create_example,
                project_path=src_project_path,
                ignored_items=",".join(["README.md"]),
                overwrite=False,
            )
        elif language == L.golang:
            ctx.invoke(
                golang.create_example_console,
                project_path=src_project_path,
                ignored_items=",".join(["README.md"]),
                overwrite=False,
            )
        elif language == L.cpp:
            ctx.invoke(
                cpp.create_example,
                project_path=src_project_path,
                ignored_items=",".join(["README.md"]),
                overwrite=False,
            )
        elif language == L.kotlin:
            # Kotlin project is complex, so it doesn't create an example automatically.
            # Use IntelliJ IEDA to create it.
            if read_user_yes_no("Launch explorer?", "yes"):
                os.makedirs(src_project_path, exist_ok=True)
                click.launch(src_project_path, locate=True)
