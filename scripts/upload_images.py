'''
Date: 2021.06.14 15:24
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2021.06.14 15:24
'''
import os.path
import sys
from tempfile import gettempdir

import requests

headers = {
    'sec-ch-ua': '" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"',
    'accept': 'application/json, text/javascript, */*; q=0.01',
    'dnt': '1',
    'x-requested-with': 'XMLHttpRequest',
    'sec-ch-ua-mobile': '?0',
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 '
                  '(KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36',
    'accept-language': 'zh-CN,zh;q=0.9,en;q=0.8,ca;q=0.7',
}

upload_url = "https://image.kieng.cn/upload.html?type=jd"
success_list = []
tempdir = gettempdir()


def download_image(url: str) -> str:
    file = os.path.join(tempdir, os.path.basename(url))
    with open(file, 'wb') as fp:
        fp.write(requests.get(url, headers=headers).content)
    return file


def upload_image(path: str):
    if path.startswith('http'):
        path = download_image(path)

    if os.path.isfile(path):
        files = {'image': open(path, 'rb')}
        res = requests.post(upload_url, headers=headers, files=files).json()

        if res['code'] == 200:
            success_list.append(res['data']['url'])
        else:
            success_list.append(res['msg'])


def script_upload_images():
    for path in sys.argv[1:]:
        upload_image(path)

    print('Upload Success:')
    print('\n'.join(success_list))


if __name__ == '__main__':
    # download_image('https://image.taoguba.com.cn/img/2021/06/14/kn2lc4c7ysj6.png_760w.png')

    script_upload_images()
