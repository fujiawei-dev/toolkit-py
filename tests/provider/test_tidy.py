import tempfile

from PIL import Image

from toolkit.provider.tidy import is_binary_file


def test_is_binary_file():
    text_file = tempfile.mktemp()
    with open(text_file, "w") as f:
        f.write("test")

    assert not is_binary_file(text_file)

    binary_file = tempfile.mktemp()
    Image.new("RGB", (100, 100)).save(binary_file, "JPEG")

    assert is_binary_file(binary_file)
