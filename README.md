# Toolkit-Py

[![PyPI](https://img.shields.io/pypi/v/toolkit-py)](https://pypi.org/project/toolkit-py/)
[![License](https://img.shields.io/pypi/l/toolkit-py)](https://github.com/fujiawei-dev/toolkit-py/blob/master/LICENSE)

[![Python Test](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-test.yml/badge.svg)](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-test.yml)
[![Python Publish](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-publish.yml/badge.svg)](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-publish.yml)

## Table of Contents

- [Toolkit-Py](#toolkit-py)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Automatically unzip files recursively](#automatically-unzip-files-recursively)
    - [Recursively change the encoding of text files in the current folder](#recursively-change-the-encoding-of-text-files-in-the-current-folder)
    - [Generates HTTP User-Agent header](#generates-http-user-agent-header)
    - [Change to other faster open source mirror sites](#change-to-other-faster-open-source-mirror-sites)
    - [Create basic project scaffold](#create-basic-project-scaffold)
    - [Upload pictures to public image hosting server](#upload-pictures-to-public-image-hosting-server)
    - [Create or display the configuration of commonly used software](#create-or-display-the-configuration-of-commonly-used-software)

## Installation

```shell
pip install -U toolkit-py
```

```shell
pip install -U toolkit-py -i https://pypi.douban.com/simple
```

## Usage

### Remove the ui ads of YoudaoNote

```shell
$ ucmd youdao
```

### Automatically unzip files recursively

[解压嵌套加密压缩文件](unified_command/README.md#解压嵌套加密压缩文件)

### Recursively change the encoding of text files in the current folder

```shell
$ cen
c05_mbr.asm: GB2312 -> utf-8
c17_core.asm: GB2312 -> utf-8
c17_mbr.asm: GB2312 -> utf-8
nasmide.ini: ascii -> utf-8
```

### Generates HTTP User-Agent header

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

### Change to other faster open source mirror sites

```shell
$ cfm
Usage: cfm [OPTIONS] COMMAND [ARGS]...

Options:
  --help  Show this message and exit.

Commands:
  python (py)       Change pypi & conda source minors.
  raspberrypi (pi)  Change Raspberry Pi OS source minors.
  ubuntu (ubuntu)   Change ubuntu/ubuntu-port source minors.
```

### Create basic project scaffold

```shell
$ cps
Usage: cps [OPTIONS] COMMAND [ARGS]...

Options:
  --help  Show this message and exit.

Commands:
  base                            Create basic project scaffold.
  c                               Create C/C++ project scaffold.
  clean (clean,clear,release,rm)  Remove all example files for release.
  golang (go)                     Create Golang project scaffold.
  notes                           Create notes project scaffold.
  python (py)                     Create Python project scaffold.
  qt5                             Create Qt5 project scaffold.
```

### Upload pictures to public image hosting server

```shell
$ upsfortypora test.png
Upload Success:
http://dd-static.jd.com/ddimg/jfs/t1/132543/17/21538/145549/61fa87f9E883b9b32/f23efa3a806cab76.jpg
```

### Create or display the configuration of commonly used software

```shell
$ ccf
Usage: ccf [OPTIONS] COMMAND [ARGS]...

Options:
  --help  Show this message and exit.

Commands:
  alias            Generate aliases for powershell or bash configuration
                   files.
  clash            Display configuration file of clash.
  hosts            Display configuration file of hosts.
  notes (nn)       Create a new note for hugo.
  powershell (ps)  Create or display configuration files about PowerShell.
  python (py)      Create or display configuration files about Python.
```

```shell
 $ ccf py -h
 Usage: ccf py [OPTIONS]

  Create or display configuration files about Python.

Options:
  -m, --method INTEGER     1 -> .pypirc
  -r, --read-only BOOLEAN  Read only or create configuration files.
  -h, --help               Show this message and exit.
```
