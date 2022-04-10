"""
Date: 2022.04.10 16:21
Description: Omit
LastEditors: Rustle Karl
LastEditTime: 2022.04.10 16:21
"""
import os

from .common import writer


def vim(read_only=True):
    conf = os.path.join(os.path.expanduser("~"), ".vimrc")

    content = """\
set nocompatible " 不与 Vi 兼容
syntax on " 打开语法高亮
set showmode " 在底部显示，当前处于命令模式还是插入模式
set showcmd " 命令模式下，在底部显示，当前键入的指令
set mouse=a " 支持使用鼠标
set encoding=utf-8 " 使用 utf-8 编码
set t_Co=256 " 启用256色
filetype indent on " 开启文件类型检查，并且载入与该类型对应的缩进规则
set autoindent " 按下回车键后，下一行的缩进会自动跟上一行的缩进保持一致
set shiftwidth=2 " 设置缩进宽度为 2 个空格 
set tabstop=2 " 设置 Tab 键宽度为 2 个空格 
set expandtab " 自动将 Tab 转为空格
set softtabstop=2 " Tab 转为多少个空格
set number " 显示行号
set cursorline " 光标所在的当前行高亮
set textwidth=80 " 设置行宽，即一行显示多少个字符
set wrap " 自动折行，即太长的行分成几行显示
set linebreak " 不在单词内部折行
set wrapmargin=2 " 指定折行处与编辑窗口的右边缘之间空出的字符数
set scrolloff=5 " 垂直滚动时，光标距离顶部/底部的位置行数
set sidescrolloff=15 " 水平滚动时，光标距离行首或行尾的位置行数
set laststatus=2 " 是否显示状态栏。0 表示不显示，1 表示只在多窗口时显示，2 表示显示
set ruler " 在状态栏显示光标的当前位置
set showmatch " 光标遇到圆括号、方括号、大括号时，自动高亮对应的另一个圆括号、方括号和大括号
set hlsearch " 搜索时，高亮显示匹配结果
set incsearch " 输入搜索模式时，每输入一个字符，就自动跳到第一个匹配的结果
set ignorecase " 搜索时忽略大小写
set smartcase " 如果同时打开了ignorecase，那么对于只有一个大写字母的搜索词，将大小写敏感；其他情况都是大小写不敏感
set history=1000 " 记住多少次历史操作
set wildmenu
set wildmode=longest:list,full " 命令模式下，底部操作指令按下 Tab 键自动补全。第一次按下 Tab，会显示所有匹配的操作指令的清单；第二次按下 Tab，会依次选择各个指令
"""
    writer(conf, content=content, read_only=read_only)
