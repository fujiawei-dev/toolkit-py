"""
Date: 2022.02.02 20:24
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 20:24
"""
import os.path
from dataclasses import dataclass
from tempfile import gettempdir

import requests

headers = {
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
        return f"success={self.success}, message={self.message}, remote_url={self.remote_url}"


def download_image(url: str) -> str:
    file = os.path.join(gettempdir(), os.path.basename(url))

    with open(file, "wb") as fp:
        fp.write(requests.get(url, headers=headers).content)

    return file
