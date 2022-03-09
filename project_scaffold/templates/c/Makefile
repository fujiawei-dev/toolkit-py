{{HASHTAG_COMMENTS}}

.PHONY: ;
#.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.IGNORE: clean;            # ignore all errors, keep going

APP_VERSION = 0.0.1
APP_NAME = {{APP_NAME}}

UNIFIED_NAME := $(APP_NAME)
INSTALL_PATH = /usr/local/bin/
ifeq ($(OS), Windows_NT)
UNIFIED_NAME := $(UNIFIED_NAME).exe
INSTALL_PATH = c:/developer/bin/
endif

CMAKE_BUILD_PREFIX = cmake-build
CMAKE_OPTIONS = -G "Unix Makefiles"
MAKE_OPTIONS = -j 64

#VCPKG_PATH = C:/Developer/vcpkg
#CMAKE_OPTIONS += -DCMAKE_TOOLCHAIN_FILE=$(VCPKG_PATH)/scripts/buildsystems/vcpkg.cmake

all: run

uninstall:
	-(rm $(INSTALL_PATH)$(UNIFIED_NAME))

install: build-release
	copy $(CMAKE_BUILD_PREFIX)-release/$(UNIFIED_NAME) $(INSTALL_PATH)

reinstall: uninstall install

build-debug:
	cd $(CMAKE_BUILD_PREFIX)-debug; cmake $(CMAKE_OPTIONS) -DCMAKE_BUILD_TYPE=Debug ..; make $(MAKE_OPTIONS)

build-release:
	cd $(CMAKE_BUILD_PREFIX)-release; cmake $(CMAKE_OPTIONS) -DCMAKE_BUILD_TYPE=Release ..; make $(MAKE_OPTIONS)

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

tag:
	git tag v$(APP_VERSION)
	git push origin v$(APP_VERSION)
