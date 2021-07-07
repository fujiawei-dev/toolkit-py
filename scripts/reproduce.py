'''Automatically create some common files.'''

from os.path import isfile
from datetime import datetime

# README.md
tmpl_readme = '''\
# README
'''

# upload.sh
tmpl_upload = '''\
git init
git add LICENSE
git commit -m "add: LICENSE"

git add go.mod
git commit -m "add: mod"

git add README.md
git commit -m "add: README.md"

git push -u origin master
'''

# .gitignore
tmpl_gitignore = '''\
# 自定义文件
upload.sh
w_*
replace
*ttf
fonts
cc
.sync_folder.json

# IDE 文件
.idea
.vscode
.vscode-test/
.vscodeignore

# Dart
.dart_tool/
.packages
build/
doc/api/

# Go 相关
vendor
go.sum
*.exe

# Python 相关
venv
__pycache__
build
dist
*egg-info

# JavaScript / TypeScript
out
node_modules
*.vsix
*.lock
.yarnrc

# 日志文件
*.log
logs

# 存储文件
uploads
storage
*.db
testdata
_gsdata_

# Dropbox
*.paper

.dockerignore
.gitignore\
'''

tmpl_license = '''\
The MIT License (MIT)

Copyright (c) 2021 Rustle Karl

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.\
'''.replace('2021', str(datetime.now().year))

tmpl_vscodeignore = '''\
.vscode/**
.vscode-test/**
out/test/**
src/**
.gitignore
.yarnrc
vsc-extension-quickstart.md
**/tsconfig.json
**/.eslintrc.json
**/*.map
**/*.ts
'''

files = ['.gitignore', 'LICENSE', 'upload.sh', 'README.md']
tmpls = [tmpl_gitignore, tmpl_license, tmpl_upload, tmpl_readme]

# files.append('.vscodeignore')
# tmpls.append(tmpl_vscodeignore)


def script_rpd():
    for file, tmpl in zip(files, tmpls):
        if isfile(file):
            continue
        print(tmpl, file=open(file, 'w', encoding='utf-8'))

    print('OK')
