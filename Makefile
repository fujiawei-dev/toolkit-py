.PHONY: clean coverage dist help install lint

define BROWSER_PYSCRIPT
import os, webbrowser, sys

from urllib.request import pathname2url

webbrowser.open("file://" + pathname2url(os.path.abspath(sys.argv[1])))
endef
export BROWSER_PYSCRIPT

BROWSER := python -c "$$BROWSER_PYSCRIPT"

define PRINT_HELP_PYSCRIPT
import re, sys

for line in sys.stdin:
	match = re.match(r'^([a-zA-Z_-]+):.*?## (.*)$$', line)
	if match:
		target, help = match.groups()
		print("%-20s %s" % (target, help))
endef
export PRINT_HELP_PYSCRIPT

help: ## print help
	@python -c "$$PRINT_HELP_PYSCRIPT" < $(MAKEFILE_LIST)

version: ## print version
	@python setup.py version

pip: ## install dependencies
	pip install -i https://pypi.douban.com/simple -r requirements.txt
	pip install -i https://pypi.douban.com/simple -r requirements-dev.txt

black: ## format code
	@black .

flake8: ## check style
	@flake8 .

lint: black flake8 ## check style

test: ## run tests quickly with the default Python
	python setup.py test
	-@pytest tests

coverage: ## view coverage report
	$(BROWSER) htmlcov/index.html

.IGNORE: clean-pyc

clean: clean-build clean-pyc clean-test ## remove all build, test, coverage and Python artifacts

clean-build: ## remove build artifacts
	rm -fr build/
	rm -fr dist/
	rm -fr .eggs/
	rm -fr UNKNOWN.egg-info/
	-find . -name '*.egg-info' -exec rm -fr {} +
	-find . -name '*.egg' -exec rm -f {} +

clean-pyc: ## remove Python file artifacts
	find . -name '*.pyc' -exec rm -f {} +
	find . -name '*.pyo' -exec rm -f {} +
	find . -name '*~' -exec rm -f {} +
	find . -name '__pycache__' -exec rm -fr {} +

clean-test: ## remove test and coverage artifacts
	rm -fr .tox/
	rm -f .coverage
	rm -f coverage.xml
	rm -fr htmlcov/
	rm -fr .pytest_cache

release: dist ## package and upload a release
	twine upload dist/*

dist: clean ## builds source and wheel package
	python setup.py sdist
	python setup.py bdist_wheel
	ls -l dist

install: clean ## install the package to the active Python's site-packages
	python setup.py install

act-install: ## install act
	curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
	cp ./bin/act /usr/local/bin

act: ## run act
	act -v
	act

env: ## print environment
	@echo $(shell cat .env)

init: ## initialize a new project
	git init
	versioneer install
	pre-commit install
