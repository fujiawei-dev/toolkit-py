import os
import re
from pathlib import Path

import click

from toolkit.provider.image_hosting import download_image

IMAGE_URL_PATTERN = re.compile(r"!\[.*?\]\((https?://.*?)\)")


def replace_image_url_from_file(file_path: str, images_path: str) -> bool:
    with open(file_path, "r", encoding="utf-8") as f:
        content = f.read()

    urls = IMAGE_URL_PATTERN.findall(content)

    if urls and not os.path.exists(images_path):
        os.makedirs(images_path, exist_ok=True)

    for url in urls:
        image_path = os.path.join(images_path, os.path.basename(url))
        download_image(url, image_path)
        content = content.replace(
            url,
            os.path.relpath(
                image_path,
                os.path.dirname(file_path),
            ).replace("\\", "/"),
        )

    if urls:
        with open(file_path, "w", encoding="utf-8", newline="\n") as f:
            f.write(content)

    return bool(urls)


@click.command(
    help="Replace online image url to local image path.",
    context_settings={"ignore_unknown_options": True},
)
@click.option(
    "--workspace-path",
    type=click.STRING,
    default=os.getcwd(),
    help="The path of the workspace.",
)
@click.argument(
    "article-path",
    type=click.Path(exists=True),
    default=os.getcwd(),
    required=False,
    nargs=1,
)
def offline_images(workspace_path, article_path):
    for item in Path(article_path).rglob("*.md"):
        if item.is_file():
            if replace_image_url_from_file(
                item.as_posix(),
                os.path.join(
                    workspace_path,
                    "assets",
                    "images",
                    os.path.relpath(os.path.splitext(item)[0], workspace_path),
                ),
            ):
                click.echo(f"Replace image url from {item}")
