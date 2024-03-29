APP_NAME = {{ project_slug.snake_case }}

QT_VERSION = {{ qt_version }}
QT_TOOL_VERSION = {{ qt_tool_version }}
QT_COMPILE_VERSION = {{ qt_compile_version }}

ifeq ($(OS), Windows_NT)
	SHELL = pwsh
	SHELLFLAGS = -Command
	QT_PATH = C:/Qt/Qt${QT_VERSION}
	EXECUTABLE ?= ${APP_NAME}.exe
else
	EXECUTABLE ?= ${APP_NAME}
endif

CMAKE_PREFIX_PATH = ${QT_PATH}/${QT_VERSION}/${QT_COMPILE_VERSION}

WIN_DEPLOY_QT = $$Env:PATH="${CMAKE_PREFIX_PATH}/bin;\
							${QT_PATH}/Tools/${QT_TOOL_VERSION}/bin;"+$$Env:Path; \
							$$Env:QML2_IMPORT_PATH="${CMAKE_PREFIX_PATH}/qml"; \
							$$Env:QT_DEBUG_PLUGINS=1; \
							cd ${DIST_DIR}; \
							windeployqt --release --qmldir "${CMAKE_PREFIX_PATH}/qml" ${EXECUTABLE}

BIN_DIR = bin

DIST_DIR = dist

BUILD_DIR = cmake-build-release-${QT_TOOL_VERSION}

dist: dist-copy dist-deploy-qt dist-installer

dist-clean:
	-rm -r ${DIST_DIR}
	-mkdir ${DIST_DIR}
	-rm -r ${BIN_DIR}
	-mkdir ${BIN_DIR}

dist-copy: dist-clean
	cp ${BUILD_DIR}/${EXECUTABLE} ${DIST_DIR}/${EXECUTABLE}
	cp -r lib/*.dll ${DIST_DIR}

dist-deploy-qt:
	pwsh -Command '${WIN_DEPLOY_QT}'
	mv ${DIST_DIR}/${EXECUTABLE} ${BIN_DIR}/${EXECUTABLE}

dist-installer:
	enigmavbconsole installer/enigma_setup.evb
	iscc installer/inno_setup.iss
