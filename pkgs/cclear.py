'''
Date: 2020-12-17 15:17:56
LastEditors: Rustle Karl
LastEditTime: 2021-01-01 18:21:45
'''

import os
from glob import glob


def cclear():
    if not os.path.isdir('cc'):
        os.mkdir('cc')
    for old in glob('*.go'):
        clear_go_cmts(old, 'cc/'+old)
    for old in glob('*.py'):
        clear_py_cmts(old, 'cc/'+old, False)


def clear_go_cmts(old: str, new: str):
    '''Clear golang comments.'''
    old_fp = open(old, 'r', encoding='utf-8')
    new_fp = open(new, 'w', encoding='utf-8')
    line = old_fp.readline()
    while line:
        if line.lstrip().startswith('//'):
            continue
        if '//' in line:
            new_fp.write(line[:line.index('//')])
            continue
        new_fp.write(line)
        line = old_fp.readline()
    old_fp.close()
    new_fp.close()


def clear_py_cmts(old: str, new: str, strict=True):
    '''Clear python comments.'''
    old_fp = open(old, 'r', encoding='utf-8')
    new_fp = open(new, 'w', encoding='utf-8')
    flag = 0
    line = old_fp.readline()
    while line:
        if line.lstrip().startswith("'''") or line.lstrip().startswith('"""'):
            flag += 1
            continue
        if line.rstrip().endswith("'''") or line.rstrip().endswith('"""'):
            flag = 0
            continue
        if flag % 2 == 1:
            continue
        if '#' in line and strict:
            new_fp.write(line[:line.index('#')])
            continue
        new_fp.write(line)
        line = old_fp.readline()
    old_fp.close()
    new_fp.close()
