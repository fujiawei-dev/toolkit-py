# Toolkit-Py

[![License](https://img.shields.io/pypi/l/toolkit-py)](https://github.com/fujiawei-dev/toolkit-py/blob/main/LICENSE)
[![Latest Version](https://img.shields.io/pypi/v/toolkit-py)](https://pypi.org/project/toolkit-py/)
[![Supported Python versions](https://img.shields.io/pypi/pyversions/toolkit-py)](https://pypi.python.org/pypi/toolkit-py)

[![Python Test](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-test.yml/badge.svg)](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-test.yml)
[![Python Publish](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-publish.yml/badge.svg)](https://github.com/fujiawei-dev/toolkit-py/actions/workflows/python-publish.yml)

> Personal toolkit implemented by Python.

## Installation

```shell
pip install -U toolkit-py
```

If you are in China, you can use the following command to install the latest version:

```shell
pip install -U toolkit-py -i https://pypi.douban.com/simple
```

Install it from source code:

```shell
pip install git+https://github.com/fujiawei-dev/toolkit-py.git@main
```

## Usage

```shell
toolkit --help
```

## Thanks

![JetBrains Logo (Main) logo](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)

## Examples

### Generates HTTP User-Agent header

```shell
gua
```

```
Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/15.14986
```

```shell
gua -n chrome
```

```
Mozilla/5.0 (X11; Linux i686) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3489.10 Safari/537.36
```

```shell
gua -o android
```

```
Mozilla/5.0 (Linux; Android 8.1; Huawei P20 Lite Build/OPR3.170623.008) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3297.48 Mobile Safari/537.36
```

```shell
gua -n safari -o ios
```

```
Mozilla/5.0 (iPhone; CPU iPhone OS 9_3_3 like Mac OS X) AppleWebKit/602.2.14 (KHTML, like Gecko) Version/9.0 Mobile/13G34 Safari/602.2
```

### Upload pictures to public image hosting server

```shell
upsfortypora "/path/to/file.png"
```

```
Upload Success:
http://dd-static.jd.com/ddimg/jfs/t1/132543/17/21538/145549/61fa87f9E883b9b32/f23efa3a806cab76.jpg
```
