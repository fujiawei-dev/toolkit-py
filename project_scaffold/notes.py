'''
Date: 2022.02.03 12:37
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.03 12:37
'''
from .common import create_common_files


def notes():
    create_common_files([
        'quickstart',  # 安装/卸载/基础用法
        'docs',  # 基础语法/知识/原理
        'libraries',  # 库
        'libraries/standard',  # 标准库
        'libraries/tripartite',  # 第三方库
        'assets/images',  # 图片
        'src',  # 源码示例
        'src/libraries/standard',  # 标准库源码示例
        'src/libraries/tripartite',  # 第三方库源码示例
    ])
