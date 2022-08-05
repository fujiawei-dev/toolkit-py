import tempfile
from pathlib import Path

from PIL import Image

from toolkit.provider import tidy


def test_is_binary_file():
    text_file = tempfile.mktemp()
    with open(text_file, "w") as f:
        f.write("test")

    assert not tidy.is_binary_file(text_file)

    binary_file = tempfile.mktemp()
    Image.new("RGB", (100, 100)).save(binary_file, "JPEG")

    assert tidy.is_binary_file(binary_file)


def test_replace_all():
    old = (
        "Duis ea kasd sed rebum elit invidunt amet ipsum velit"
        " invidunt tempor lorem facilisi mazim."
    )
    words = ["ea", "sd", "lo"]
    new = tidy.replace_all(old, words)
    for word in words:
        assert word not in new


def test_clean_characters():
    pairs = [
        ("OK", "OK\n"),
        ("\nOK\t\n", "OK\n"),
        ("\n\t\nOK\t\n\n", "OK\n"),
    ]

    for pair in pairs:
        assert tidy.clean_characters(pair[0]) == pair[1]


def test_change_encoding():
    content = "Python3 translate()方法| 菜鸟教程"
    with tempfile.TemporaryDirectory() as tmpdir:
        src = Path(tmpdir) / "test.txt"
        src.write_text(content, encoding="GB2312")
        tidy.change_encoding(src)
        assert src.read_text(encoding="utf-8") == content
