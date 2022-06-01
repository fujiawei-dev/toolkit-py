{{HASHTAG_COMMENTS}}

APP_NAME = {{APP_NAME}}

QT_VERSION = 5.12.6

ifeq ($(OS), Windows_NT)
	QT_PATH = C:/Qt
    QT_COMPILE_VERSION = mingw73_32
    QT_TOOL_VERSION = mingw730_32
	EXECUTABLE ?= ${APP_NAME}.exe
else
	EXECUTABLE ?= ${APP_NAME}
endif

CMAKE_PREFIX_PATH = ${QT_PATH}/Qt${QT_VERSION}/${QT_VERSION}/${QT_COMPILE_VERSION}

DEPLOY_CMD = $$Env:PATH="${CMAKE_PREFIX_PATH}/bin;\
						${QT_PATH}/Qt${QT_VERSION}/Tools/${QT_TOOL_VERSION}/bin;"+$$Env:Path; \
						$$Env:QML2_IMPORT_PATH="${CMAKE_PREFIX_PATH}/qml"; \
						$$Env:QT_DEBUG_PLUGINS=1; \
						cd ${DIST_DIR}; \
						windeployqt --release --qmldir "${CMAKE_PREFIX_PATH}/qml" ${EXECUTABLE}

BIN_DIR = bin
DIST_DIR = dist
BUILD_DIR = cmake-build-release

info:
	@echo $(SHELL)

init:
	git init
	git add .
	git commit -m "first commit"

clean: clean-build clean-dist

clean-build:
	-rm -r ${BUILD_DIR}
	-mkdir ${BUILD_DIR}

clean-dist:
	-(rm -r ${DIST_DIR})
	-(mkdir ${DIST_DIR})

clean-bin:
	-rm -r ${BIN_DIR}
	-mkdir ${BIN_DIR}

pack: pack-deployqt pack-setup

pack-copy: clean-dist clean-bin
	cp ${BUILD_DIR}/${EXECUTABLE} ${DIST_DIR}/${EXECUTABLE}
	cp -r assets ${DIST_DIR}
	cp -r lib/*.dll ${DIST_DIR}
	cp pack/assets/settings.ini ${BIN_DIR}/settings.ini

pack-deployqt: SHELL:=pwsh
pack-deployqt: pack-copy
	${DEPLOY_CMD}
	mv ${DIST_DIR}/${EXECUTABLE} ${BIN_DIR}/${EXECUTABLE}

pack-setup:
	enigmavbconsole pack/enigma_setup.evb
	iscc pack/inno_setup.iss
	rm ${BIN_DIR}/${EXECUTABLE}

pack-compress:
	-(rm ${BIN_DIR}/${APP_NAME}.7z)
	7z a ${BIN_DIR}/${APP_NAME}.7z -p321 ${BIN_DIR}/*

# ONLY FOR TESTING!

USER = root
HOST = 192.168.1.87

USER_HOST = "$(USER)@$(HOST)"
SSH_CMD = ssh $(USER_HOST)

UPLOAD_PATH = /root/project/storage/uploads/$(APP_NAME)

pack-upload: pack
	$(SSH_CMD) "mkdir -p $(UPLOAD_PATH)"
	scp bin/*_Setup.exe "$(USER)@$(HOST):$(UPLOAD_PATH)"
