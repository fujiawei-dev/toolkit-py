'''
Date: 2020-09-21 23:48:26
Description: This module is for generating random, valid web navigator's User-Agent HTTP headers.
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 15:01
'''

'''
Command Line Usage

```shell
$ gua
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/15.14986

$ gua -n chrome
Mozilla/5.0 (X11; Linux i686) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3489.10 Safari/537.36

$ gua -o android
Mozilla/5.0 (Linux; Android 8.1; Huawei P20 Lite Build/OPR3.170623.008) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3297.48 Mobile Safari/537.36

$ gua -n safari -o ios
Mozilla/5.0 (iPhone; CPU iPhone OS 9_3_3 like Mac OS X) AppleWebKit/602.2.14 (KHTML, like Gecko) Version/9.0 Mobile/13G34 Safari/602.2
```
'''

from argparse import ArgumentParser

from .generate import generate_user_agent


def command_gua():
    parser = ArgumentParser(
            usage='%(prog)s [options] usage',
            description='Generates User-Agent HTTP header',
    )

    parser.add_argument('-o', '--os',
                        help='limit list of os for generation, possible values:\
                        "win", "linux", "mac", "android", "ios", "all"')

    parser.add_argument('-n', '--browser',
                        help='limit list of browser engines for generation, possible values:\
                        "chrome", "firefox", "edge", "safari", "opera", "all"')

    parser.add_argument('-d', '--platform', help='possible values:\
                        "desktop", "smartphone", "all"')

    opts = parser.parse_args()

    gua = generate_user_agent(os=opts.os, browser=opts.browser, platform=opts.platform)

    print(gua)
