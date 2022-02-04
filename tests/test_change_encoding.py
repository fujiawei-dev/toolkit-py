'''
Date: 2022.02.04 11:09
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.02.04 11:09
'''
import os

from change_encoding.command import change_encoding

cwd = os.path.abspath(os.path.dirname(__file__))


# The file name cannot be prefixed with test, otherwise pytest will try to open it
# and report an error if there is a problem with the encoding
def test_change_encoding():
    change_encoding(os.path.join(cwd, 'data/gbktext.txt'), os.path.join(cwd, 'data/utf8text.txt'), encoding='gbk')
    change_encoding(os.path.join(cwd, 'data/utf8text.txt'))
