import shutil
from pathlib import Path
from typing import Union


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
