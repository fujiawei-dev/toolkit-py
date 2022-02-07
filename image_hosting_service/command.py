"""
Date: 2022.02.02 18:14
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 18:14
"""
import os.path
import sys

from .common import UploadResult, download_image
from .jd import upload as upload_to_jd


def upload_image(path: str) -> UploadResult:
    if path.startswith("http"):
        path = download_image(path)

    if os.path.isfile(path):
        return upload_to_jd(path)

    return UploadResult(message=f"file {path} not is file")


def command_ups_for_typora():
    success_list = []

    for path in sys.argv[1:]:
        result = upload_image(path)

        if result.success:
            success_list.append(result.remote_url)
        else:
            success_list.append(result.message)

    print("Upload Success:")
    print("\n".join(success_list))
