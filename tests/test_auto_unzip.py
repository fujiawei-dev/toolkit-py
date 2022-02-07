"""
Date: 2022.02.07 23:03
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.07 23:03
"""
import shutil

from unified_command.auto_unzip import Unzipper


def test_auto_unzip():
    unzipper = Unzipper()
    test_path = unzipper.create_7z_files_for_test()
    assert unzipper.run(test_path)
    shutil.rmtree(test_path)
