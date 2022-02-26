{{MAKEFILE_HEADER}}

.PHONY: ;
#.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.IGNORE: clean;            # ignore all errors, keep going

APP_VERSION = 0.0.1
APP_NAME = {{APP_NAME}}

UNIFIED_NAME := $(APP_NAME)
ifeq ($(OS), Windows_NT)
UNIFIED_NAME := $(UNIFIED_NAME).exe
endif

CMAKE_BUILD_DIR = cmake-build-debug

all: run

build: clean
	cd $(CMAKE_BUILD_DIR); cmake -G "Unix Makefiles" ..; make

run: build
	cd $(CMAKE_BUILD_DIR); ./$(UNIFIED_NAME)

clean:
	-(mkdir $(CMAKE_BUILD_DIR))
	cd $(CMAKE_BUILD_DIR); rm -r *
