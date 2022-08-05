import os
import sys

from requests import Session

from .common import UploadResult, download_image
from .jd import upload as upload_to_jd


def upload_image(path: str, session: Session = Session()) -> UploadResult:
    if path.startswith("http"):
        path = download_image(path)

    if os.path.isfile(path):
        return upload_to_jd(path, session)

    return UploadResult(message=f"file {path} not is file")


def upload_images(paths: list) -> list[UploadResult]:
    return [upload_image(path) for path in paths]


def ups_for_typora_command():
    results = upload_images(sys.argv[1:])
    print("Upload Success:")
    print(
        "\n".join(
            result.remote_url if result.success else result.message
            for result in results
        )
    )
