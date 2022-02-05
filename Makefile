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

VERSION = 0.2.5
PACKAGE = toolkit-py

all: test install clean

dep:
	pip install -r requirements.txt

# 打包
build:
	python setup.py sdist
	python setup.py bdist_wheel

# 安装
install: build
	pip install --force-reinstall --no-deps dist/$(PACKAGE)-$(VERSION).tar.gz

upload: build
	twine upload dist/$(PACKAGE)-$(VERSION).tar.gz

test:
	pytest
	rm -r .pytest_cache

#清理编译中间文件
clean:
	rm -r build
	rm -r dist
	rm -r *egg-info

tag:
	git tag v$(VERSION)
	git push origin v$(VERSION)
