.PHONY: ;
.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

.IGNORE: dep clean test;            # ignore all errors, keep going

PACKAGE = toolkit-py

VERSION := $(shell python -c "from unified_command.version import __version__; print(__version__, end='')")

all: format reinstall test

format:
	pip install -U black
	black .

version:
	echo $(VERSION)

dep:
	pip install -i https://pypi.douban.com/simple -r requirements.txt

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
	pytest tests
	rm -r .pytest_cache

clean:
	rm -r build
	rm -r dist
	rm -r *egg-info
	rm -r $(PACKAGE)-$(VERSION)
	rm -r .pytest_cache

tag:
	git tag v$(VERSION)
	git push origin v$(VERSION)
