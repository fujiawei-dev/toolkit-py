import tempfile
from pathlib import Path

from toolkit.fs.dir import copy_items_in_directory


def test_copy_items_in_directory():
    with tempfile.TemporaryDirectory() as tmpdir:
        src = Path(tmpdir) / "src"
        dst = Path(tmpdir) / "dst"

        src.mkdir()
        dst.mkdir()
        (src / "a").mkdir()
        (src / "a" / "a").touch()

        copy_items_in_directory(src, dst)

        assert (dst / "a" / "a").exists()
        assert not (dst / "src").exists()
