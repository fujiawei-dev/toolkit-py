{{HASHTAG_COMMENTS}}

.PHONY: ;
.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

.IGNORE: dep clean tests;            # ignore all errors, keep going

SHELL = bash

PYPI_PACKAGE = {{PYPI_PACKAGE}}

APP_VERSION := $(shell python -c "from {{PYTHON_MODULE}}.version import __version__; print(__version__, end='')")

APP_NAME_VERSION := $(PYPI_PACKAGE)_$(APP_VERSION)

all: format reinstall tests

info:
	@echo $(SHELL)

init:
	git init
	git add .
	git commit -m "first commit"

format:
	pip install -U black
	black .

version:
	echo $(APP_VERSION)

dep:
	pip install -r requirements.txt

build:
	python setup.py sdist
	#python setup.py bdist_wheel
	#python -m build

uninstall:
	pip uninstall -y $(PYPI_PACKAGE)

install: uninstall build
	pip install --force-reinstall --no-deps dist/$(PYPI_PACKAGE)-$(APP_VERSION).tar.gz

reinstall: install clean

test:
	pytest
	rm -r .pytest_cache

clean:
	rm -r build
	rm -r dist
	rm -r *egg-info
	rm -r .pytest_cache

tag:
	git tag v$(APP_VERSION)
	git push origin v$(APP_VERSION)

build-onefile:
	pyinstaller --onefile -y --clean -n $(APP_NAME_VERSION) main.py
	-(rm $(APP_NAME_VERSION).spec)
	-(rm -r build)
