"""
Date: 2022.02.02 19:56
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 19:56
"""
import base64
import json

import lxml.html
import requests

from .common import UploadResult, headers

session = requests.Session()

session.headers = headers


def upload(path) -> UploadResult:
    result = UploadResult()

    response = session.post(
        "https://imio.jd.com/uploadfile/file/post.do",
        {
            "appId": "im.customer",
            "clientType": "comet",
            "s": base64.b64encode(open(path, "rb").read()),
        },
    )

    if response.status_code == 200:
        html: lxml.html.HtmlElement = lxml.html.fromstring(response.content)
        body = json.loads(html.find("body").text)

        if body["code"] == 0:
            result.success = True
            result.remote_url = body["path"]

        result.message = body["desc"]

    return result
