from tests.provider.test_unzip_data import (
    assert_compressed_files,
    create_compressed_files,
)
from toolkit.provider.unzip import Unzipper, delete_empty_directories


def test_create_compressed_files():
    create_compressed_files("data")


def test_unzipper():
    unzipper = Unzipper()

    path = "data"
    move_to = "data/unzipped"

    path = create_compressed_files(
        path=path,
        passwords=unzipper.passwords[1:10],
    )

    unzipper.run(path, move_to)

    assert_compressed_files(move_to, path)

    delete_empty_directories(path)
