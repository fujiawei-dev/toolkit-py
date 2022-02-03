'''
Date: 2022.02.03 11:22
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.03 11:22
'''
import os

from image_hosting_service.command import upload_image

cwd = os.path.abspath(os.path.dirname(__file__))


def test_upload_local_image():
    result = upload_image(os.path.join(cwd, 'data/test.jpg'))

    if not result.success:
        raise ValueError('Upload image failed: %s' % result.message)

    print('Image remote url: %s' % result.remote_url)


def test_upload_remote_image():
    result = upload_image('https://www.baidu.com/img/flexible/logo/pc/peak-result.png')

    if not result.success:
        raise ValueError('Upload image failed: %s' % result.message)

    print('Image remote url: %s' % result.remote_url)
