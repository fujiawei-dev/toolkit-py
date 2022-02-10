{{MAKEFILE_HEADER}}

.PHONY: ;
.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

.IGNORE: dep clean tests;            # ignore all errors, keep going

PYPI_PACKAGE = {{PYPI_PACKAGE}}

APP_VERSION := $(shell python -c "from unified_command.version import __version__; print(__version__, end='')")

all: format reinstall tests

format:
	pip install -U black
	black .

version:
	echo $(VERSION)

dep:
	pip install -r requirements.txt

build:
	python setup.py sdist
	#python setup.py bdist_wheel
	#python -m build

uninstall:
	pip uninstall -y $(PACKAGE)

install: uninstall build
	pip install --force-reinstall --no-deps dist/$(PACKAGE)-$(VERSION).tar.gz

reinstall: install clean

test:
	pytest
	rm -r .pytest_cache

clean:
	rm -r build
	rm -r dist
	rm -r *egg-info

tag:
	git tag v$(VERSION)
	git push origin v$(VERSION)
