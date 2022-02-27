"""
Date: 2022.02.07 15:09
Description: Automatically unzip files recursively.
LastEditors: Rustle Karl
LastEditTime: 2022.02.08 18:46:44
"""
import contextlib
import os.path
import pickle
import re
import shutil
import subprocess
import time
from collections import defaultdict
from pathlib import Path
from queue import LifoQueue, Queue
from shutil import which
from threading import Thread
from typing import Dict, List, Tuple

PASSWORDS_DIR = os.path.normpath(os.path.expanduser("~/.config/.passwords"))
PASSWORDS_FILE = os.path.normpath(os.path.join(PASSWORDS_DIR, "customize.txt"))

INCLUDED_SUFFIXES = {".7z", ".zip", ".rar", ".7zz"}

EXCLUDE_SUFFIXES = {
    ".downloading",
    ".jpeg",
    ".jpg",
    ".mp3",
    ".mp4",
    ".png",
    ".torrent",
    ".pickle",
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
    print("[debug] src=%s dst=%s password=%r" % (src, dst, password))

    cmd = ["7z", "x", src, "-o" + dst, "-p" + password, "-aou"]

    if not is_7zip_utf8_encoding(src):
        cmd.append("-mcp=936")

    print(f"[debug] run: {' '.join(cmd)}")

    success = subprocess.run(cmd, capture_output=True).returncode == 0
    if not success:
        shutil.rmtree(dst, ignore_errors=True)

    return success


def _cmd_7zip_list(src) -> bytes:
    """获取列表"""
    p = subprocess.run(["7z", "l", "-ba", src, "-p"], capture_output=True, timeout=1)
    return p.stdout or p.stderr


def is_7zip_utf8_encoding(src) -> bool:
    try:
        content = _cmd_7zip_list(src)
        return "�" not in content.decode("utf-8")
    except subprocess.TimeoutExpired:
        return True


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
        print(f"[debug] password={repr(password)}?")
        if _cmd_7zip_decompress(src, dst, password):
            return True, password, dst

    return False, "", dst


def _need_continue(src, i=0, j=0) -> Tuple[int, int]:
    """判断是否继续解压"""
    if not os.path.exists(src):
        return 1, 0

    if os.path.isfile(src):
        return 0, 1

    for file in os.listdir(src):
        suffix = os.path.splitext(file)[1].lower()
        file = os.path.join(src, file)
        if os.path.isfile(file):
            if suffix in INCLUDED_SUFFIXES:
                return 0, 1
            if suffix in EXCLUDE_SUFFIXES:
                i += 1  # 已经解压的可信度
            else:
                j += 1  # 继续解压的可信度
        elif os.path.isdir(file):
            i, j = _need_continue(file, i, j)

    return i, j


def normalize_segment_zips(file, suffixes: list) -> list:
    print(f"[debug] {file}: {suffixes}")

    files = []
    for suffix in suffixes:
        src = os.path.normpath(file + suffix)
        suffix = ".7z.%03d" % int(re.sub(r"[^\d]", "0", suffix[suffix.rfind(".") :]))
        dst = os.path.normpath(file + suffix)
        if src != dst and not os.path.exists(dst):
            os.renames(src, dst)
        files.append(dst)
    files.sort()
    return files


class Unzipper(object):
    # 密码字典
    _passwords_dictionary: List[str] = [""]
    _passwords_loaded = False

    _failed_items = Queue()  # 失败列队
    _successful_items = LifoQueue()  # 成功列队
    _successful_items_dict = dict()  # 保存成功项目，以便从记录中恢复

    # 有可能是分卷压缩文件
    _segment_zips: Dict[str, Queue] = defaultdict(Queue)

    def __init__(self):
        self._loads_passwords()

    def _loads_passwords(self):
        """载入密码字典"""
        if self._passwords_loaded:
            return

        if not os.path.exists(PASSWORDS_DIR):
            os.makedirs(PASSWORDS_DIR)

        if os.path.isfile(PASSWORDS_FILE):
            with open(PASSWORDS_FILE, encoding="utf-8") as fp:
                self._passwords_dictionary.extend(fp.read().splitlines())
        else:
            with open(PASSWORDS_FILE, "wb") as fp:
                fp.write(b"\n")

        for file in os.listdir(PASSWORDS_DIR):
            if os.path.isfile(file) and not is_equal_path(file, PASSWORDS_FILE):
                with open(file, encoding="utf-8") as fp:
                    self._passwords_dictionary.extend(fp.read().splitlines())

        self._passwords_loaded = True

    def _cmd_7zip(self, src, dst, parent) -> Tuple[bool, str, str]:
        """单文件解压"""
        success, password, dst = self._successful_items_dict.get(
            src
        ) or cmd_7zip_decompress(src, dst, self._passwords_dictionary)

        if success:
            self._successful_items_dict[src] = (True, password, dst)
            self._successful_items.put((os.path.relpath(src, parent), password))
        else:
            self._failed_items.put((os.path.relpath(src, parent)))

        return success, password, dst

    def _run_recursive(self, src, move_to, parent, segment=False) -> bool:
        """嵌套压缩解压"""
        if os.path.isfile(src):
            suffix = os.path.splitext(src)[1].lower()

            if suffix in EXCLUDE_SUFFIXES:
                return True

            if suffix in DELETE_SUFFIXES:
                Path(src).unlink(missing_ok=True)
                return True

            # 已经确认是分卷压缩
            if suffix not in INCLUDED_SUFFIXES and not segment:
                p = Path(src)

                if len(p.suffixes) > 1 and " " not in p.suffixes[0]:  # 多个后缀可能是分卷压缩
                    stem = p.stem
                    pos = stem.find(".")

                    if pos != -1:
                        stem = stem[:pos]

                    k = (p.parent / stem).as_posix()
                    v = p.as_posix()[len(k) :]
                    self._segment_zips[k].put(v)

                    return True

            if suffix == "":  # 无后缀名的情况自动添加后缀
                _src = src + ".7z"
                if not os.path.exists(_src):
                    os.renames(src, _src)
                src = _src

            print("[debug] unzip: %s" % src)
            success, _, src = self._cmd_7zip(src, "", parent)

            if success and os.path.exists(src):
                i, j = _need_continue(src)
                if j > 0 and (i < j or i < 8):
                    print("[debug] continue unzip: %s" % src)
                    return self._run_recursive(src, move_to, parent)

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

            return success

        elif not is_equal_path(src, move_to) and os.path.isdir(src):
            # 多线程处理

            ts: List[Thread] = []

            for file in os.listdir(src):
                t = Thread(
                    target=self._run_recursive,
                    args=(os.path.join(src, file), move_to, parent),
                    daemon=True,
                )
                t.start()
                ts.append(t)

            for t in ts:
                t.join()

            # FIXME: How to get return value?

        return False

    def run(self, src, move_to="") -> bool:
        begin = time.time()  # 计时器

        if move_to == "" or move_to == src:
            move_to = os.path.dirname(src) if os.path.isfile(src) else src
            move_to = os.path.join(move_to, "Unzip")

        os.makedirs(move_to, exist_ok=True)

        self._run_recursive(src, move_to, src)

        for k in list(self._segment_zips.keys()):
            v = self._segment_zips[k]
            files = normalize_segment_zips(k, list(v.queue))
            if self._run_recursive(files[0], move_to, src, segment=True):
                for file in files:
                    Path(file).unlink(missing_ok=True)
                self._segment_zips[k] = Queue()

        if not self._successful_items.empty():
            print("\n=========== INFO ============")

        while not self._successful_items.empty():
            file, password = self._successful_items.get()
            print(f"{file}: {password}")

            # 删除已经解压成功的源文件
            file = os.path.join(src, file)
            if os.path.isfile(file):
                Path(file).unlink(missing_ok=True)
                with contextlib.suppress(OSError):
                    os.removedirs(os.path.dirname(file))
            elif os.path.isdir(file):
                os.removedirs(file)

        if self._segment_zips:
            print("\n=========== WARNING ============")

            for k, v in self._segment_zips.items():
                if not v.empty():
                    print(k, ":", list(v.queue))

        no_failure = True

        if not self._failed_items.empty():
            no_failure = False
            print("\n=========== ERROR ============")

        while not self._failed_items.empty():
            print(self._failed_items.get())

        print(f"\n\ncost {time.time() - begin:.02f}s")

        return no_failure

    def run_with_history(self, src, move_to="", history_file="history.pickle") -> bool:
        if os.path.exists(history_file):
            self._successful_items_dict = pickle.load(file=open(history_file, "rb"))

        try:
            return self.run(src, move_to)
        finally:
            pickle.dump(self._successful_items_dict, file=open(history_file, "wb"))

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


if __name__ == "__main__":
    unzipper = Unzipper()
    unzipper.run_with_history("s:/cache/done/test")
