{{HASHTAG_COMMENTS}}

.PHONY: ;
#.SILENT: ;               # no need for @
#.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.IGNORE: clean;            # ignore all errors, keep going

APP_NAME = {{APP_NAME}}

UNIFIED_NAME := $(APP_NAME)
ifeq ($(OS), Windows_NT)
UNIFIED_NAME := $(UNIFIED_NAME).exe
endif

# Qt default installation directory
QT_PATH = C:/Qt/Qt5.12.6
QT_VERSION = 5.12.6
QT_COMPILE_VERSION = mingw73_32
QT_TOOL_VERSION = mingw730_32

CMAKE_PREFIX_PATH = ${QT_PATH}/${QT_VERSION}/${QT_COMPILE_VERSION}

CMAKE_BUILD_PREFIX = cmake-build
RELEASE_PACKAGE = cmake-build-package

all: clean

info:
	@echo $(SHELL)

init:
	git init
	git add .
	git commit -m "first commit"

clean:
	-(mkdir $(CMAKE_BUILD_PREFIX)-debug)
	-(cd $(CMAKE_BUILD_PREFIX)-debug; rm -r *)

	-(mkdir $(CMAKE_BUILD_PREFIX)-release)
	-(cd $(CMAKE_BUILD_PREFIX)-release; rm -r *)

release: release-copy release-deployqt

release-copy:
	-(rm -r $(RELEASE_PACKAGE))
	-(mkdir $(RELEASE_PACKAGE))
	cp $(CMAKE_BUILD_PREFIX)-release/$(UNIFIED_NAME) $(RELEASE_PACKAGE)/$(UNIFIED_NAME)
	cp -r assets $(RELEASE_PACKAGE)
	cp -r lib/*.dll $(RELEASE_PACKAGE)

release-deployqt: SHELL:=pwsh   # This is setting the shell for b only, but you should disable ONESHELL.
release-deployqt: # $(eval SHELL:=pwsh)
	@echo $(SHELL)
	$$Env:PATH="${CMAKE_PREFIX_PATH}/bin;${QT_PATH}/Tools/${QT_TOOL_VERSION}/bin;"+$$Env:Path; $$Env:QML2_IMPORT_PATH="${CMAKE_PREFIX_PATH}/qml"; cd $(RELEASE_PACKAGE); windeployqt --release --qmldir "${CMAKE_PREFIX_PATH}/qml" ${UNIFIED_NAME}

release-run:
	@echo $(SHELL)
	cd $(RELEASE_PACKAGE); ./${UNIFIED_NAME}

release-clean:
	-(rm -r $(RELEASE_PACKAGE)/logs)
	-(rm $(RELEASE_PACKAGE)/settings.ini)
