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

CMAKE_BUILD_PREFIX = cmake-build

all: run

build-debug:
	cd $(CMAKE_BUILD_PREFIX)-debug; cmake -DCMAKE_BUILD_TYPE=Debug -G "Unix Makefiles" ..; make

build-release:
	cd $(CMAKE_BUILD_PREFIX)-release; cmake -DCMAKE_BUILD_TYPE=Release -G "Unix Makefiles" ..; make

build: clean build-release

run-debug: build-debug
	cd $(CMAKE_BUILD_PREFIX)-debug; ./$(UNIFIED_NAME)

run-release: build-release
	cd $(CMAKE_BUILD_PREFIX)-release; ./$(UNIFIED_NAME)

run: clean run-release

clean:
	-(mkdir $(CMAKE_BUILD_PREFIX)-debug)
	cd $(CMAKE_BUILD_PREFIX)-debug; rm -r *

	-(mkdir $(CMAKE_BUILD_PREFIX)-release)
	cd $(CMAKE_BUILD_PREFIX)-release; rm -r *
