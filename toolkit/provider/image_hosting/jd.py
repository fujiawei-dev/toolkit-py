import base64
import json

import lxml.html
from requests import Session

from toolkit.provider.image_hosting.common import UploadResult


def upload(path: str, session: Session = Session()) -> UploadResult:
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
