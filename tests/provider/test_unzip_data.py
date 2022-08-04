import os.path
import shutil
import tempfile
from random import choice

from toolkit.provider.unzip import compress

PASSWORDS = ["password", "test", "root", "admin"]

DELETABLE_FILES = ["link.url", "web.txt"]

IMAGE_JPG_FILES = [f"{i:03d}.jpg" for i in range(1, 15)]

COMPRESSED_FILES = {
    "普通压缩文件.7z": IMAGE_JPG_FILES,
    "无后缀压缩文件": IMAGE_JPG_FILES,
    "嵌套压缩文件.7z": [
        {
            "嵌套内部压缩文件1.7z": IMAGE_JPG_FILES,
            "嵌套内部压缩文件2.7z": IMAGE_JPG_FILES,
            "嵌套内部压缩文件3.7z": IMAGE_JPG_FILES,
        },
        *DELETABLE_FILES,
    ],
}


def assert_compressed_files(
    unzip_path: str = "",
    parent_path: str = "",
    compressed_files: dict[str, list] = None,
):
    compressed_files = compressed_files or COMPRESSED_FILES

    for key, value in compressed_files.items():
        folder = os.path.splitext(key)[0]
        if parent_path and os.path.exists(parent_path):
            assert not os.path.exists(os.path.join(parent_path, key))
            assert not os.path.exists(os.path.join(parent_path, folder))
        folder = os.path.join(unzip_path, folder)
        for item in value:
            if isinstance(value, dict):
                assert_compressed_files(folder, compressed_files=item)
            elif isinstance(value, str):
                assert os.path.exists(os.path.join(folder, item))


def create_compressed_files(
    path: str = "",
    parent_temp_dir: str = None,
    passwords: list[str] = None,
    compressed_files: dict[str, list] = None,
) -> str:
    path = path or tempfile.mkdtemp()

    if not os.path.exists(path):
        os.mkdir(path)

    passwords = passwords or PASSWORDS
    compressed_files = compressed_files or COMPRESSED_FILES

    for key, value in compressed_files.items():
        temp_dir = tempfile.mkdtemp(dir=parent_temp_dir)
        temp_dir = os.path.join(temp_dir, os.path.splitext(key)[0])
        os.mkdir(temp_dir)

        for item in value:
            if isinstance(item, str):
                open(os.path.join(temp_dir, item), "wb").write(b"x" * 1024 * 1024)
            elif isinstance(item, dict):
                create_compressed_files(temp_dir, parent_temp_dir, passwords, item)

        dst = os.path.join(path, key)
        compress(temp_dir, dst, choice(passwords))
        if os.path.splitext(key)[1] == "":
            os.renames(dst + ".7z", dst)

        shutil.rmtree(temp_dir)

    return path
