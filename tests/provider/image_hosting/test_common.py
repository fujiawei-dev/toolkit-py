import os.path
import sys

from toolkit.provider import image_hosting
from toolkit.provider.image_hosting import common

IMAGE_URLS = [
    "https://www.baidu.com/img/bd_logo1.png",
    "https://img-blog.csdnimg.cn/img_convert/19c566899007ac0b4cf84b8ed4debe29.png",
]


def test_download_image():
    for url in IMAGE_URLS:
        assert os.path.exists(common.download_image(url))


def test_upload_image():
    for url in IMAGE_URLS:
        assert image_hosting.upload_image(url).success


def test_upload_images():
    for result in image_hosting.upload_images(IMAGE_URLS):
        assert result.success


def test_ups_for_typora_command():
    sys.argv = ["upsfortypora"] + IMAGE_URLS
    image_hosting.ups_for_typora_command()
