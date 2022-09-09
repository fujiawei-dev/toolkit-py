import os.path
from pathlib import Path
from typing import NamedTuple, Union

import click
import yaml


@click.command(
    help="Extract the titles of all articles from the folder as the table of contents."
)
@click.option(
    "--folder-path",
    "-p",
    type=click.Path(exists=True, file_okay=False, dir_okay=True),
    default=Path.cwd(),
    required=True,
    help="The path of the articles folder.",
)
@click.option(
    "--output-path",
    "-o",
    type=click.Path(exists=False, file_okay=True, dir_okay=False),
    default=Path.cwd() / "toc.md",
    required=True,
    help="The path of the output file.",
)
def extract_toc_from_folder(
    folder_path: Union[str, Path],
    output_path: Union[str, Path],
):
    """从文件夹中提取全部笔记的标题作为总目录"""

    class Article(NamedTuple):
        uri: str
        title: str

    titles: list[Article] = []

    for article_path in Path(folder_path).glob("**/*.md"):
        if article_path.open().read(3) != "---":
            continue

        with article_path.open(encoding="utf-8") as fp:
            content = ""
            started = False

            for line in fp:
                if started and line.startswith("---"):
                    break

                if not started and line.startswith("---"):
                    started = True
                elif started:
                    content += line

            metadata = yaml.safe_load(content)
            titles.append(
                Article(
                    os.path.relpath(article_path, os.path.dirname(output_path)),
                    metadata["title"],
                )
            )

    with Path(output_path).open("w", encoding="utf-8") as fp:
        for article in titles:
            fp.write(f"- [{article.title}]({article.uri})\n".replace("\\", "/"))
