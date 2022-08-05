"""
Date: 2020-09-21 23:48:26
Description: This module is for generating random, valid web navigator's User-Agent HTTP headers.
LastEditors: Rustle Karl
LastEditTime: 2022.02.02 15:01
"""
from argparse import ArgumentParser

from .generate import generate_user_agent


def generate_user_agent_command():
    parser = ArgumentParser(
        usage="%(prog)s [options] usage",
        description="Generates User-Agent HTTP header",
    )

    parser.add_argument(
        "-o",
        "--os",
        help='limit list of os for generation, possible values:\
                        "win", "linux", "mac", "android", "ios", "all"',
    )

    parser.add_argument(
        "-n",
        "--browser",
        help='limit list of browser engines for generation, possible values:\
                        "chrome", "firefox", "edge", "safari", "opera", "all"',
    )

    parser.add_argument(
        "-d",
        "--platform",
        help='possible values:\
                        "desktop", "smartphone", "all"',
    )

    opts = parser.parse_args()

    print(
        generate_user_agent(
            os=opts.os,
            browser=opts.browser,
            platform=opts.platform,
        )
    )
