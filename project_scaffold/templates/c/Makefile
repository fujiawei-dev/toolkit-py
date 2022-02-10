{{MAKEFILE_HEADER}}

.PHONY: ;
#.SILENT: ;			   # no need for @
.ONESHELL: ;			 # recipes execute in same shell
.NOTPARALLEL: ;		  # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.IGNORE: clean;			# ignore all errors, keep going

ENTRY_POINT = {{ENTRY_POINT}}

APP_VERSION = 0.0.1
APP_NAME = {{APP_NAME}}

INSTALL_PATH = /usr/local/bin/

UNIFIED_NAME := $(APP_NAME)
ifeq ($(OS), Windows_NT)
UNIFIED_NAME := $(UNIFIED_NAME).exe
INSTALL_PATH = c:/developer/bin/
endif

all: reinstall

build:
	gcc -o $(UNIFIED_NAME) $(ENTRY_POINT)

uninstall:
	-(rm $(INSTALL_PATH)$(UNIFIED_NAME))

install:
	copy $(UNIFIED_NAME) $(INSTALL_PATH)
	rm $(UNIFIED_NAME)

reinstall: uninstall install

tag:
	git tag v$(VERSION)
	git push origin v$(VERSION)
