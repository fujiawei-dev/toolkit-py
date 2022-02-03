'''
Date: 2022.02.03 9:17
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.03 9:17
'''
import os

import click


def writer(conf, content='', read_only=True, official=''):
    if not read_only and content != '':
        os.makedirs(os.path.dirname(conf), exist_ok=True)
        with open(conf, 'w', encoding='utf8') as fp:
            fp.write(content)

    click.echo(official)
    click.echo(conf)
