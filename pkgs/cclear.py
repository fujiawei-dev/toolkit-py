'''
Date: 2020-12-17 15:17:56
LastEditors: Rustle Karl
LastEditTime: 2020-12-18 08:21:16
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
    while (l := old_fp.readline()):
        if l.lstrip().startswith('//'):
            continue
        if '//' in l:
            new_fp.write(l[:l.index('//')])
            continue
        new_fp.write(l)
    old_fp.close()
    new_fp.close()


def clear_py_cmts(old: str, new: str, strict=True):
    '''Clear python comments.'''
    old_fp = open(old, 'r', encoding='utf-8')
    new_fp = open(new, 'w', encoding='utf-8')
    flag = 0
    while (l := old_fp.readline()):
        if l.lstrip().startswith("'''") or l.lstrip().startswith('"""'):
            flag += 1
            continue
        if l.rstrip().endswith("'''") or l.rstrip().endswith('"""'):
            flag = 0
            continue
        if flag % 2 == 1:
            continue
        if '#' in l and strict:
            new_fp.write(l[:l.index('#')])
            continue
        new_fp.write(l)
    old_fp.close()
    new_fp.close()
