import os.path
from dataclasses import dataclass
from tempfile import gettempdir

import requests

HEADERS = {
    "accept": "application/json, text/javascript, */*; q=0.01",
    "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "
    "(KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36",
    "accept-language": "zh-CN,zh;q=0.9,en;q=0.8,ca;q=0.7",
}


@dataclass
class UploadResult(object):
    success: bool = False
    message: str = ""
    remote_url: str = ""

    def __str__(self):
        return (
            f"success={self.success}, "
            f"message={self.message}, "
            f"remote_url={self.remote_url}"
        )


def download_image(url: str) -> str:
    file = os.path.join(gettempdir(), os.path.basename(url))

    with open(file, "wb") as fp:
        fp.write(requests.get(url, headers=HEADERS).content)

    return file
