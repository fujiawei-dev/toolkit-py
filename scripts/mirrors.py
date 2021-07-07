'''Change mirror source.'''

import sys
import os

from argparse import ArgumentParser


# ============= PyPi =============
pypi_mirrors = '''\
[global]
index-url=http://mirrors.aliyun.com/pypi/simple/

[install]
trusted-host=mirrors.aliyun.com
'''

# https://mirrors.ustc.edu.cn/pypi/web/simple/
# https://mirrors.tuna.tsinghua.edu.cn/pypi/web/simple/


def change_pypi():
    user = os.path.expanduser('~')

    if sys.platform == 'win32':
        cfg = user + "/appdata/roaming/pip/pip.ini"
    elif sys.platform == 'linux' or sys.platform == 'darwin':
        cfg = user + "/.pip/pip.conf"

    os.makedirs(os.path.dirname(cfg), exist_ok=True)
    with open(cfg, 'w', encoding='utf-8') as fp:
        fp.write(pypi_mirrors)


# ============= Conda =============
conda_mirrors = '''\
channels:
  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main/
  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/
  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/conda-forge/
  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/pytorch/
  - https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/pytorch3d/
ssl_verify: false

auto_activate_base: false
'''


def change_conda():
    user = os.path.expanduser('~')
    cfg = user + "/.condarc"

    with open(cfg, 'w', encoding='utf-8') as fp:
        fp.write(conda_mirrors)


# ============================
def change_source(who):

    if who is None or who == "all":
        change_pypi()
        change_conda()
    elif who == "py":
        change_pypi()
    elif who == "conda":
        change_conda()

    print("Completed")


def script_chs():
    parser = ArgumentParser(
        usage='%(prog)s [options] usage',
        description='Change mirror source',
    )

    parser.add_argument(
        '-w', '--who', help='optional values: "py", "conda", "all"')
    opts = parser.parse_args()
    change_source(who=opts.who)
