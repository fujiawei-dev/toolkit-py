
APP_NAME = {{ project_slug.snake_case }}

QT_VERSION = {{ qt_version }}
QT_COMPILE_VERSION = {{ qt_compile_version }}
QT_TOOL_VERSION = {{ qt_tool_version }}

ifeq ($(OS), Windows_NT)
	SHELL = pwsh
	SHELLFLAGS = -Command
	EXECUTABLE ?= ${APP_NAME}.exe
	INSTALL_DIR = C:/Developer/bin
	QT_PATH = C:/Qt/Qt${QT_VERSION}

CMAKE_PREFIX_PATH = ${QT_PATH}/${QT_VERSION}/${QT_COMPILE_VERSION}

WINDEPLOYQT = $$Env:PATH="${CMAKE_PREFIX_PATH}/bin;\
						${QT_PATH}/Qt${QT_VERSION}/Tools/${QT_TOOL_VERSION}/bin;"+$$Env:Path; \
						$$Env:QML2_IMPORT_PATH="${CMAKE_PREFIX_PATH}/qml"; \
						$$Env:QT_DEBUG_PLUGINS=1; \
						cd ${DIST_DIR}; \
						windeployqt --release --qmldir "${CMAKE_PREFIX_PATH}/qml" ${EXECUTABLE}

BIN_DIR = bin

DIST_DIR = dist

BUILD_DIR = cmake-build-release-${QT_TOOL_VERSION}

clean: clean-build clean-dist clean-bin

clean-build:
	-rm -r ${BUILD_DIR}
	-mkdir ${BUILD_DIR}

clean-dist:
	-(rm -r ${DIST_DIR})
	-(mkdir ${DIST_DIR})

clean-bin:
	-rm -r ${BIN_DIR}
	-mkdir ${BIN_DIR}

pack: pack-copy pack-deployqt pack-setup

pack-copy: clean-dist clean-bin
	cp ${BUILD_DIR}/${EXECUTABLE} ${DIST_DIR}/${EXECUTABLE}
	cp -r assets ${DIST_DIR}
	cp -r lib/*.dll ${DIST_DIR}

pack-deployqt:
	${DEPLOY_CMD}
	mv ${DIST_DIR}/${EXECUTABLE} ${BIN_DIR}/${EXECUTABLE}

pack-setup:
	enigmavbconsole pack/enigma_setup.evb
	iscc pack/inno_setup.iss
	rm ${BIN_DIR}/${EXECUTABLE}
