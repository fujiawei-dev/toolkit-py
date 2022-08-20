import contextlib
import os.path
import pickle
import re
import shlex
import shutil
import subprocess
import time
from collections import defaultdict
from pathlib import Path
from queue import LifoQueue, Queue
from shutil import which
from subprocess import PIPE
from threading import Thread
from typing import Dict, Iterable, List, Tuple, Union

import click
import yaml

from toolkit.config.runtime import EDITOR
from toolkit.logger import logging

log = logging.getLogger(__name__)

DEFAULT_PASSWORDS_DIR = Path.home() / ".config" / ".passwords"
DEFAULT_PASSWORDS_DIR.mkdir(parents=True, exist_ok=True)
DEFAULT_PASSWORDS_FILE = DEFAULT_PASSWORDS_DIR / "customize.txt"
DEFAULT_PASSWORDS_FILE.touch(exist_ok=True)

NORMAL_FILE_SUFFIXES = {".jpeg", ".jpg", ".png", ".mp3"}
COMPRESSED_FILE_SUFFIXES = {".7z", ".zip", ".rar", ".7zz"}
DELETABLE_FILE_SUFFIXES = {".txt", ".url", ".html", ".gif", ".mp4", ".docx"}
IGNORE_FILE_SUFFIXES = {".pickle"}


def delete_empty_directories(path: Union[str, Path]):
    for root, dirs, files in os.walk(path, topdown=False):
        if not files and not os.listdir(root):
            log.warning(f"delete empty directory {root}")
            os.rmdir(root)


def list_compressed(src: str) -> bytes:
    args = ["7z", "l", "-ba", src, "-p"]
    proc = subprocess.run(args, capture_output=True, timeout=1)
    return proc.stdout or proc.stderr


def is_utf8_encoding(src: str) -> bool:
    try:
        content = list_compressed(src)
        return "�" not in content.decode("utf-8")
    except subprocess.TimeoutExpired:
        return True


def decompress(src: str, dst: str, password: str) -> bool:
    log.debug(f"compress {src} to {dst} with {password!r}")

    args = ["7z", "x", src, "-o" + dst, "-p" + password, "-aou"]

    if not is_utf8_encoding(src):
        args.append("-mcp=936")

    log.debug(shlex.join(args))

    success = False

    timeout = os.path.getsize(src) // (1 << 24)
    timeout = 8 if password == "" else max(timeout, 5)

    try:
        success = (
            subprocess.call(
                args,
                timeout=timeout,
                stdout=PIPE,
                stderr=PIPE,
            )
            == 0
        )
    except subprocess.TimeoutExpired:
        log.warning(f"{src} decompress with {password!r} timeout {timeout}s")

    if not success:
        shutil.rmtree(dst, ignore_errors=True)

    return success


def compress(src: Union[str, Iterable], dst: str, password: str) -> bool:
    if len(src) == 0:
        return True

    log.debug(f"compress {src} to {dst} with {password!r}")

    args = ["7z", "a", dst, "-p" + password]

    if isinstance(src, str):
        args.append(src)
    elif isinstance(src, Iterable):
        args.extend(src)

    log.debug(shlex.join(args))

    if not (success := subprocess.call(args) == 0):
        if os.path.exists(dst):
            os.remove(dst)

    return success


def brute_force_decompress(
    src: str,
    dst: str = "",
    passwords: list = None,
) -> Tuple[bool, str, str]:
    if not which("7z"):
        raise FileNotFoundError("7z not found.")

    dst = dst or os.path.splitext(src)[0]
    passwords = passwords or []

    for password in passwords:
        log.debug(f"try to decompress {src} with {password!r}")
        if decompress(src, dst, password):
            log.info(f"{src} found password {password!r}")
            return True, password, dst

    log.warning(f"{src} not found password")
    return False, "", dst


def continue_decompress(
    src: str,
    normal_files: int = 0,
    compressed_files: int = 0,
) -> Tuple[int, int]:
    if os.path.exists(src):
        if os.path.isfile(src):
            compressed_files += 1

        elif os.path.isdir(src):
            for item in os.listdir(src):
                suffix = os.path.splitext(item)[1].lower()
                item = os.path.join(src, item)

                if os.path.isfile(item):
                    if suffix in COMPRESSED_FILE_SUFFIXES:
                        compressed_files += 99
                        break

                    if suffix in NORMAL_FILE_SUFFIXES:
                        normal_files += 1
                    else:
                        compressed_files += 1

                elif os.path.isdir(item):
                    normal_files, compressed_files = continue_decompress(
                        item, normal_files, compressed_files
                    )

    return normal_files, compressed_files


def normalize_segment_zips(file, suffixes: list) -> list:
    log.debug(f"{file}: {suffixes}")
    files = []
    for suffix in suffixes:
        src = os.path.normpath(file + suffix)
        suffix = ".7z.%03d" % int(re.sub(r"\D", "0", suffix[suffix.rfind(".") :]))
        dst = os.path.normpath(file + suffix)
        if src != dst and not os.path.exists(dst):
            os.renames(src, dst)
        files.append(dst)
    files.sort()
    return files


class Unzipper(object):
    def __init__(self):
        self.loaded: bool = False
        self.passwords: List[str] = [""]

        self.failed_items: Queue = Queue()
        self.successful_items: Queue = Queue()
        self.successful_items_dict: Dict[str, Tuple] = dict()

        self.segment_zips: Dict[str, Queue] = defaultdict(Queue)

        self.load_passwords()

    def load_passwords(self, default_only: bool = True, max_size: int = 100):
        if not self.loaded:
            if not DEFAULT_PASSWORDS_DIR.exists():
                DEFAULT_PASSWORDS_DIR.mkdir(parents=True)

            if not DEFAULT_PASSWORDS_FILE.exists():
                DEFAULT_PASSWORDS_FILE.touch()
            else:
                self.passwords.extend(
                    DEFAULT_PASSWORDS_FILE.read_text(encoding="utf-8").splitlines()
                )

            if not default_only:
                for file in DEFAULT_PASSWORDS_DIR.iterdir():
                    if file.is_file() and file != DEFAULT_PASSWORDS_FILE:
                        self.passwords.extend(
                            file.read_text(encoding="utf-8").splitlines()
                        )

                self.passwords = self.passwords[:max_size]

            self.loaded = True

    def decompress(self, src: str, dst: str, parent: str) -> Tuple[bool, str, str]:
        item = self.successful_items_dict.get(src)
        item = item or brute_force_decompress(src, dst, self.passwords)
        success, password, dst = item

        if success:
            self.successful_items_dict[src] = (True, password, dst)
            self.successful_items.put((os.path.relpath(src, parent), password))
        else:
            self.failed_items.put((os.path.relpath(src, parent)))

        return success, password, dst

    def decompress_recursively(
        self,
        src: str,
        move_to: str,
        parent: str,
        segment=False,
    ) -> bool:
        if os.path.isfile(src):
            suffix = os.path.splitext(src)[1].lower()

            if suffix in IGNORE_FILE_SUFFIXES:
                return False

            if suffix in NORMAL_FILE_SUFFIXES:
                return True

            if suffix in DELETABLE_FILE_SUFFIXES:
                return os.unlink(src) is None

            if suffix not in COMPRESSED_FILE_SUFFIXES and not segment:
                p = Path(src)

                # 可能是分卷压缩文件
                if len(p.suffixes) > 1 and " " not in p.suffixes[0]:
                    stem = p.stem
                    if (pos := stem.find(".")) != -1:
                        stem = stem[:pos]
                    k = (p.parent / stem).as_posix()
                    v = p.as_posix()[len(k) :]
                    self.segment_zips[k].put(v)
                    return True

            # 添加后缀
            if suffix == "":
                new_src = src + time.strftime("_%Y%m%d%H%M%S") + ".7z"
                os.renames(src, new_src)
                src = new_src

            success, _, src = self.decompress(src, "", parent)

            if success and os.path.exists(src):
                i, j = continue_decompress(src)
                if j > 0 and (i < j or i < 8):
                    return self.decompress_recursively(src, move_to, parent)

                # 是文件夹则移动
                for item in Path(src).iterdir():
                    if item.is_dir():
                        shutil.move(item, move_to)

                try:
                    # 删除可能已经空了的文件夹
                    os.removedirs(src)
                except OSError:
                    shutil.move(src, move_to)

            return success

        elif Path(src) != Path(move_to) and os.path.isdir(src):
            log.info('working on "%s"', src)

            workers = LifoQueue(maxsize=8)

            for item in Path(src).iterdir():
                worker = Thread(
                    target=self.decompress_recursively,
                    args=(item.as_posix(), move_to, parent),
                    daemon=False,
                )

                if workers.full():
                    workers.get().join()

                worker.start()
                workers.put(worker)

            while not workers.empty():
                log.info(f"{workers.qsize()} workers are working")
                workers.get().join()

            log.info(f"{src} is successful")

        return False

    def run(self, src: str, move_to: str = "") -> bool:
        begin = time.time()  # 计时器

        if move_to == "" or Path(src) == Path(move_to):
            move_to = os.path.dirname(src) if os.path.isfile(src) else src
            move_to = os.path.join(move_to, "Unzip")

        os.makedirs(move_to, exist_ok=True)

        self.decompress_recursively(src, move_to, src)

        for k in list(self.segment_zips.keys()):
            v = self.segment_zips[k]
            files = normalize_segment_zips(k, list(v.queue))
            if self.decompress_recursively(files[0], move_to, src, segment=True):
                for file in files:
                    Path(file).unlink(missing_ok=True)
                self.segment_zips[k] = Queue()

        if not self.successful_items.empty():
            log.info("=========== INFO ============")

        while not self.successful_items.empty():
            file, password = self.successful_items.get()
            log.info(f"{file}: {password}")

            # 删除解压成功的源文件
            item = Path(src) / file
            if item.is_file():
                item.unlink(missing_ok=True)
                with contextlib.suppress(OSError):
                    item.parent.rmdir()
            elif item.is_dir():
                os.removedirs(file)

        if self.segment_zips:
            log.warning("=========== WARNING ============")

            for k, v in self.segment_zips.items():
                if not v.empty():
                    log.warning(f"{k}: {v.queue}")

        no_failure = True

        if not self.failed_items.empty():
            no_failure = False
            log.error("=========== ERROR ============")

        while not self.failed_items.empty():
            log.error(self.failed_items.get())

        for item in Path(src).iterdir():
            if item.is_dir() and item != Path(move_to):
                delete_empty_directories(src)

        log.debug(f"cost {time.time() - begin:.02f}s")

        return no_failure

    def run_with_history(
        self,
        src: str,
        move_to: str = "",
        history_file: str = "history.pickle",
    ) -> bool:
        if os.path.exists(history_file):
            self.successful_items_dict = pickle.load(file=open(history_file, "rb"))

        try:
            return self.run(src, move_to)
        except KeyboardInterrupt:
            pass
        finally:
            pickle.dump(self.successful_items_dict, file=open(history_file, "wb"))


@click.command(help="Automatically unzip files recursively.")
@click.option(
    "--src",
    type=click.Path(exists=False),
    required=False,
    default=".",
    help="The source file or directory.",
)
@click.option(
    "--move-to",
    type=click.Path(exists=False),
    required=False,
    default="",
    help="The destination directory.",
)
@click.option(
    "--disable-history",
    is_flag=True,
    help="Whether to use history file to speed up the process.",
)
@click.option(
    "--show-config",
    is_flag=True,
    help="Show the content of the default config.",
)
@click.option(
    "--open-password-file",
    is_flag=True,
    help="Open the customize password file.",
)
@click.option(
    "--open-password-dir",
    is_flag=True,
    help="Open the customize password directory.",
)
def unzip_command(
    src: str,
    move_to: str,
    disable_history: bool,
    show_config: bool,
    open_password_file: bool,
    open_password_dir: bool,
):
    if open_password_file:
        click.edit(filename=str(DEFAULT_PASSWORDS_FILE), editor=EDITOR)
        return

    if open_password_dir:
        click.launch(url=str(DEFAULT_PASSWORDS_DIR), locate=True)
        return

    if show_config:
        click.echo(
            yaml.dump(
                {
                    "default_passwords_dir": DEFAULT_PASSWORDS_DIR,
                    "default_passwords_file": DEFAULT_PASSWORDS_FILE,
                }
            )
        )
        return

    unzipper = Unzipper()

    if not disable_history:
        return unzipper.run_with_history(src, move_to)

    return unzipper.run(src, move_to)
