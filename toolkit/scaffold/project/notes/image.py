import os
import re
from pathlib import Path

import click
import requests

from toolkit.provider.image_hosting import download_image

# 网络图片正则表达式
IMAGE_URL_PATTERN = re.compile(r"!\[.*?\]\((https?://.*?)\)")

# 图片正则表达式
IMAGE_URI_PATTERN = re.compile(r"!\[.*?\]\((.*?)\)")


def replace_image_uri_from_file(file_path: str, images_path: str) -> bool:
    with open(file_path, "r", encoding="utf-8") as f:
        content = f.read()

    updated = False
    uris: list[str] = IMAGE_URI_PATTERN.findall(content)

    if uris and not os.path.exists(images_path):
        os.makedirs(images_path, exist_ok=True)

    for uri in uris:
        image_path = os.path.join(images_path, os.path.basename(uri))

        if uri.startswith("http"):  # 网络图片
            try:
                download_image(uri, image_path)
            except requests.exceptions.RequestException:
                click.secho("[x] ---------------------------------", fg="red")
                click.secho(f"[x] download image failed: {uri}", fg="red")
                click.secho(f"[x] original markdown file: {file_path}", fg="red")
                click.secho("[x] ---------------------------------", fg="red")
                return False

        else:  # 本地图片
            old_image_path = os.path.join(os.path.dirname(file_path), uri)
            old_image_path = os.path.normpath(old_image_path)
            if old_image_path != os.path.normpath(image_path):
                if not os.path.exists(old_image_path):
                    click.secho("[x] -------------------------------", fg="red")
                    click.secho(f"[x] image not exists: {old_image_path}", fg="red")
                    click.secho(f"[x] original markdown file: {file_path}", fg="red")
                    click.secho("[x] -------------------------------", fg="red")
                    return False
                else:
                    os.rename(old_image_path, image_path)

        new_uri = os.path.relpath(
            image_path,
            os.path.dirname(file_path),
        ).replace("\\", "/")

        if new_uri != uri:
            updated = True
            content = content.replace(uri, new_uri)

    if updated:
        with open(file_path, "w", encoding="utf-8", newline="\n") as f:
            f.write(content)

    return updated


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
            if replace_image_uri_from_file(
                str(item),
                os.path.join(
                    workspace_path,
                    "assets",
                    "images",
                    os.path.relpath(os.path.splitext(item)[0], workspace_path),
                ),
            ):
                click.echo(f"Replace image url from {item}")
