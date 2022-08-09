import os
import shutil
from pathlib import Path
from typing import Union

from toolkit.logger import logging

log = logging.getLogger(__name__)


def copy_items_in_directory(
    src: Union[str, Path],
    dst: Union[str, Path],
    skip_if_file_exists: bool = True,
):
    """
    Copy items in src directory to dst directory, not include itself.

    Args:
        src: Source directory.
        dst: Destination directory.
        skip_if_file_exists: If True, skip if file exists.
    """
    src, dst = Path(src), Path(dst)
    for item_src in src.iterdir():
        if item_src.is_dir():
            item_dst = dst / item_src.name
            item_dst.mkdir(parents=True, exist_ok=True)
            copy_items_in_directory(item_src, item_dst)
        elif item_src.is_file():
            item_dst = dst / item_src.name
            if not skip_if_file_exists or not item_dst.exists():
                shutil.copyfile(item_src, item_dst)


def sanitize_path(path: str):
    """Clear invalid characters in path."""
    path = path.replace("/", "、").replace("\\", "、").replace("|", "&").replace(":", "：")

    for char in ["~", '"', "?", "*", "<", ">", "{", "}"]:
        path = path.replace(char, "")

    return path


def delete_empty_folder_including_itself(path: Union[str, Path]):
    for root, dirs, files in os.walk(path, topdown=False):
        if not files and not os.listdir(root):
            log.warning(f"Deleting empty directory {root}")
            os.rmdir(root)
