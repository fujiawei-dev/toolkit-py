"""
Date: 2022.04.13 10:28
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.04.13 10:28
"""
import tempfile
from pathlib import Path

from create_config_file.notes import notes_append_header


def test_notes_append_header():
    path = tempfile.gettempdir()
    file = Path(path) / "file.txt"

    file.write_text("old_content")
    notes_append_header(file)
    print(file.read_text())
    file.unlink()

    file.write_text("---old_content")
    notes_append_header(file)
    print(file.read_text())
    file.unlink()
