"""
Date: 2022.02.07 23:03
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.07 23:03
"""
import shutil
import sys

from unified_command.auto_unzip import Unzipper


def test_auto_unzip():
    version = sys.version_info

    if version.major < 3 or version.minor < 8:
        # The missing_ok parameter was added to Path.unlink only on python 3.8.
        return

    unzipper = Unzipper()
    test_path = unzipper.create_7z_files_for_test()
    assert unzipper.run(test_path)
    shutil.rmtree(test_path)
