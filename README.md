# Toolkit-Py

![PyPI](https://img.shields.io/pypi/v/toolkit-py)

![PyPI - License](https://img.shields.io/pypi/l/toolkit-py)

[![Python application](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-app.yml/badge.svg)](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-app.yml)

[![Upload Python Package](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-publish.yml/badge.svg)](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-publish.yml)

## Installation

```shell
pip install toolkit-py
```

## Usage

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
  py      Change pypi & conda source minors.
  python  Change pypi & conda source minors.
```

### Create basic project scaffold

```shell
$ cps
Usage: cps [OPTIONS] COMMAND [ARGS]...

Options:
  --help  Show this message and exit.

Commands:
  py      Create python project scaffold.
  python  Create python project scaffold.
```

### Upload pictures to public image hosting server

```shell
$ upsfortypora test.png
Upload Success:
http://dd-static.jd.com/ddimg/jfs/t1/132543/17/21538/145549/61fa87f9E883b9b32/f23efa3a806cab76.jpg
```
