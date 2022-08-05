import os.path
import tempfile

from PIL import Image

from toolkit.provider.image_hosting import jd


def test_jd_upload():
    with tempfile.TemporaryDirectory() as tmpdir:
        file = os.path.join(tmpdir, "test.png")
        Image.new("RGB", (100, 100)).save(file)
        result = jd.upload(file)
        assert result.success
        assert result.remote_url
