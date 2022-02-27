# Unified Command 统一命令

> 将各种子命令统一到该命令之下，之前的方式每次都创建新的执行文件，一个是麻烦，每次都必须修改 setup.py 添加，另一个是可能与其他知名软件重名，污染命名空间。

```shell
$ ucmd
Usage: ucmd [OPTIONS] COMMAND [ARGS]...

Options:
  --help  Show this message and exit.

Commands:
  auto-unzip (unzip,uz)  Automatically extract files.
```

## Table of Contents

- [Unified Command 统一命令](#unified-command-统一命令)
  - [Table of Contents](#table-of-contents)
  - [解压嵌套加密压缩文件](#解压嵌套加密压缩文件)
    - [功能特性](#功能特性)
    - [显示密码字典的路径](#显示密码字典的路径)

## 解压嵌套加密压缩文件

> 有些网站为了防止资源失效，经常对一个文件多层加密，解压起来真的是浪费时间。实在是受不了了，才有了这个命令工具。

### 功能特性

- [x] 处理分卷压缩包
- [x] 删除多余空目录
- [ ] 删除成功解压后的源文件
- [ ] 密码热度排序
- [ ] 线程池

```shell
$ ucmd uz -h
Usage: ucmd uz [OPTIONS]

  Automatically extract files.

Options:
  -c, --config BOOLEAN  Show the default path of configuration file.
  -t, --test BOOLEAN    Create 7z files for test.
  -h, --help            Show this message and exit.
```

不加任何参数选项则自动解压当前文件夹。

### 显示密码字典的路径

```shell
$ ucmd uz -c
C:\Users\Admin\.config\.passwords # 文件夹，当前路径下的所有文件都认为是字典文件，全部会读取
C:\Users\Admin\.config\.passwords\customize.txt # 文件，优先级最高的字典文件
```
