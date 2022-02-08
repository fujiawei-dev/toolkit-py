"""
Date: 2022.02.07 15:09
Description: Automatically unzip files recursively.
LastEditors: Rustle Karl
LastEditTime: 2022.02.08 08:07:30
"""
import os.path
import shutil
import subprocess
import time
from queue import LifoQueue, Queue
from shutil import which
from threading import Thread
from typing import List, Tuple

PASSWORDS_DEFAULT_DIR = os.path.expanduser("~/.config/.passwords")
PASSWORDS_CUSTOMIZE_FILE = os.path.join(PASSWORDS_DEFAULT_DIR, "customize.txt")

EXCLUDE_SUFFIXES = {
    ".jpg",
    ".jpeg",
    ".png",
    ".mp4",
    ".downloading",
    ".torrent",
}

DELETE_SUFFIXES = {
    ".txt",
    ".url",
}


def is_equal_path(a, b):
    """判断路径相同"""
    return os.path.normpath(a) == os.path.normpath(b)


def _cmd_7zip_decompress(src, dst, password) -> bool:
    """解压"""
    cmd = ["7z", "x", src, "-o" + dst, "-p" + password, "-aou"]

    if not is_7zip_utf8_encoding(src):
        cmd.append("-mcp=936")

    success = subprocess.run(cmd, capture_output=True).returncode == 0
    if not success:
        shutil.rmtree(dst)

    return success


def _cmd_7zip_list(src) -> bytes:
    """获取列表"""
    p = subprocess.run(["7z", "l", "-ba", src], capture_output=True)
    return p.stdout or p.stderr


def is_7zip_utf8_encoding(src) -> bool:
    content = _cmd_7zip_list(src)
    return "�" not in content.decode("utf-8")


def _cmd_7zip_compress(src, dst, password) -> bool:
    """压缩"""
    cmd = ["7z", "a", dst, "-p" + password]
    if isinstance(src, (list, tuple)):
        cmd.extend(src)
    else:
        cmd.append(src)

    return subprocess.run(cmd, capture_output=True).returncode == 0


def cmd_7zip_decompress(
    src, dst="", passwords=None, executable="7z"
) -> Tuple[bool, str, str]:
    if not which(executable):
        raise ValueError(f"{executable} not found.")

    if not dst:
        dst = os.path.splitext(src)[0]

    if passwords is None:
        passwords = []
    elif isinstance(passwords, (str, bytes)):
        passwords = [passwords]

    for password in passwords:
        if _cmd_7zip_decompress(src, dst, password):
            return True, password, dst

    return False, "", dst


def _need_continue(src, i=0, j=0) -> Tuple[int, int]:
    """判断是否继续解压"""
    if os.path.isfile(src):
        return 0, 1

    for file in os.listdir(src):
        file = os.path.join(src, file)
        if os.path.isfile(file):
            if os.path.splitext(file)[1] in EXCLUDE_SUFFIXES:
                i += 1  # 已经解压的可信度
            else:
                j += 1  # 继续解压的可信度
        elif os.path.isdir(file):
            i, j = _need_continue(file, i, j)

    return i, j


class Unzipper(object):
    # 密码字典
    _passwords_dictionary: List[str] = [""]
    _passwords_loaded = False

    _failed_items = Queue()  # 失败列队
    _successful_items = LifoQueue()  # 成功列队

    def __init__(self):
        self._loads_passwords()

    def _loads_passwords(self):
        """载入密码字典"""
        if self._passwords_loaded:
            return

        if not os.path.exists(PASSWORDS_DEFAULT_DIR):
            os.makedirs(PASSWORDS_DEFAULT_DIR)

        if os.path.isfile(PASSWORDS_CUSTOMIZE_FILE):
            with open(PASSWORDS_CUSTOMIZE_FILE, encoding="utf-8") as fp:
                self._passwords_dictionary.extend(fp.read().splitlines())
        else:
            with open(PASSWORDS_CUSTOMIZE_FILE, "wb") as fp:
                fp.write(b"\n")

        for file in os.listdir(PASSWORDS_DEFAULT_DIR):
            if os.path.isfile(file) and not is_equal_path(
                file, PASSWORDS_CUSTOMIZE_FILE
            ):
                with open(file, encoding="utf-8") as fp:
                    self._passwords_dictionary.extend(fp.read().splitlines())

        self._passwords_loaded = True

    def _cmd_7zip(self, src, dst, parent) -> Tuple[bool, str, str]:
        """单文件解压"""
        success, password, dst = cmd_7zip_decompress(
            src, dst, self._passwords_dictionary
        )

        if success:
            self._successful_items.put((os.path.relpath(src, parent), password))
        else:
            self._failed_items.put((os.path.relpath(src, parent)))

        return success, password, dst

    def _run_recursive(self, src, move_to, parent, depth=0):
        """嵌套压缩解压"""
        if os.path.isfile(src):
            suffix = os.path.splitext(src)[1]
            if suffix in EXCLUDE_SUFFIXES:
                return
            elif suffix in DELETE_SUFFIXES:
                os.unlink(src)
                return

            if suffix == "":  # 无后缀名的情况自动添加后缀
                _src = src + ".7z"
                os.renames(src, _src)
                src = _src

            success, _, src = self._cmd_7zip(src, "", parent)

            if success:
                i, j = _need_continue(src)
                if j > 0 and (i < j or i < 8):
                    self._run_recursive(src, move_to, parent, depth + 1)
                    return

                # 如果解压出来的是文件夹组织的，则单独移动文件夹，否则一起移动
                for file in os.listdir(src):
                    file = os.path.join(src, file)
                    if os.path.isdir(file):
                        shutil.move(file, move_to)

                # 如果解压出来全是文件夹，则该目录可能已经空了，尝试删除
                try:
                    os.removedirs(src)
                except OSError:
                    shutil.move(src, move_to)

        elif not is_equal_path(src, move_to) and os.path.isdir(src):
            # 多线程处理

            ts: List[Thread] = []

            for file in os.listdir(src):
                t = Thread(
                    target=self._run_recursive,
                    args=(os.path.join(src, file), move_to, parent, depth + 1),
                )
                t.start()
                ts.append(t)

            for t in ts:
                t.join()

    def run(self, src, move_to="") -> bool:
        begin = time.time()  # 计时器

        if move_to == "" or move_to == src:
            move_to = os.path.dirname(src) if os.path.isfile(src) else src
            move_to = os.path.join(move_to, "Unzip")

        os.makedirs(move_to, exist_ok=True)

        self._run_recursive(src, move_to, src)

        while not self._successful_items.empty():
            file, password = self._successful_items.get()
            print(f"{file}: {password}")

            # 删除已经解压成功的源文件
            file = os.path.join(src, file)
            if os.path.isfile(file):
                os.unlink(file)
            elif os.path.isdir(file):
                os.removedirs(file)

        no_failure = True

        if not self._failed_items.empty():
            no_failure = False
            print("\n=========== ERROR ============")

        while not self._failed_items.empty():
            print(self._failed_items.get())

        print(f"\ncost {time.time() - begin:.02f}s")

        return no_failure

    def create_7z_files_for_test(self, path=None):
        """创建测试压缩包"""
        from random import choice

        if not path:
            import tempfile

            path = tempfile.gettempdir()

        test_path = os.path.join(path, "unzip_test")
        shutil.rmtree(test_path, ignore_errors=True)
        os.makedirs(test_path, exist_ok=True)

        # make common files
        test_jpg_files = []
        for i in range(100):
            test_jpg_file = os.path.join(test_path, f"test_file_{i:02d}.jpg")
            with open(test_jpg_file, "w") as fp:
                fp.write("test" * 1024)
            test_jpg_files.append(test_jpg_file)
            if len(self._passwords_dictionary) < 3:
                self._passwords_dictionary.append(f"test{i}")

        # make compressed files
        test_zip_files_1 = []
        for i in range(0, 100, 5):
            test_zip_file = os.path.join(test_path, f"test_zip_file_1_{i:02d}.7z")
            if _cmd_7zip_compress(
                test_jpg_files[i : i + 5],
                test_zip_file,
                choice(self._passwords_dictionary[1:]),
            ):
                test_zip_files_1.append(test_zip_file)

        for test_jpg_file in test_jpg_files:
            os.unlink(test_jpg_file)

        # make compressed files
        test_zip_files_2 = []
        for i in range(0, 15, 5):
            test_zip_file = os.path.join(test_path, f"test_zip_file_2_{i:02d}.7z")
            if _cmd_7zip_compress(
                test_zip_files_1[i : i + 5],
                test_zip_file,
                choice(self._passwords_dictionary[1:]),
            ):
                test_zip_files_2.append(test_zip_file)

        for test_zip_file_1 in test_zip_files_1[:-5]:
            os.unlink(test_zip_file_1)

        return test_path
