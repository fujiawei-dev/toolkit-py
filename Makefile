.PHONY: ;
.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

.IGNORE: dep clean;            # ignore all errors, keep going

ifeq ($(OS), Windows_NT)
SHELL := pwsh.exe
.SHELLFLAGS := -NoProfile -Command
endif

VERSION = 0.2.0
PACKAGE = toolkit-py

all: install clean

dep:
	pip install -r requirements.txt

# 打包
build:
	python setup.py sdist
	python setup.py bdist_wheel

# 安装
install: build
	pip install --force-reinstall --no-deps dist/$(PACKAGE)-$(VERSION).tar.gz

upload:
	twine upload dist/$(PACKAGE)-$(VERSION).tar.gz

#清理编译中间文件
clean:
	rm -r build
	rm -r dist
	rm -r *egg-info
