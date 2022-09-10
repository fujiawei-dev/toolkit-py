from requests import Session

from toolkit.provider.image_hosting.common import UploadResult


def upload(path: str, session: Session = Session()) -> UploadResult:
    result = UploadResult()

    response = session.post(
        "http://pic.sogou.com/pic/upload_pic.jsp",
        files={"pic_path": open(path, "rb")},
        headers={
            "Cookie": "SNUID=6FB0B40577729F3EFEC358B67763EFC7; "
            "IPLOC=CN3301; SUV=00F705FC73C3C718631BF83FDD897916; "
            "PIC_DEBUG=off; wuid=1662777407309; "
            "FUV=ac5da9c32830249b1928bd556a2a3f8f; "
            "ABTEST=0|1662777524|v1",
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0) "
            "AppleWebKit/537.36 (KHTML, like Gecko) "
            "Chrome/64.0.3282.140 Safari/537.36 Edge/15.15063",
        },
    )

    if response.status_code == 200:
        result.success = True
        result.remote_url = response.text

    return result
