import os
import tempfile

from tests.provider.test_unzip_data import (
    assert_compressed_files,
    create_compressed_files,
)
from toolkit.provider.unzip import Unzipper, delete_empty_directories


def test_create_compressed_files():
    create_compressed_files()


def test_unzipper():
    with tempfile.TemporaryDirectory() as temp_dir:
        unzipper = Unzipper()

        if len(unzipper.passwords) < 12:
            unzipper.passwords.extend("password" for _ in range(12))

        path = os.path.join(temp_dir, "path")
        move_to = os.path.join(temp_dir, "move_to")

        path = create_compressed_files(
            path=path,
            passwords=unzipper.passwords[1:10],
        )

        unzipper.run(path, move_to)

        assert_compressed_files(move_to, path)

        delete_empty_directories(path)
