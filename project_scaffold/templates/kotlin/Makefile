APP_NAME_VERSION := $(shell cd build/libs; ls *.jar)
APP_NAME := $(shell cd build/libs; ls *.jar | sed -e 's/-.*-all//g')

echo:
	@echo $(APP_NAME_VERSION)
	@echo $(APP_NAME)

build-jar:
	./gradlew.bat shadowJar --status

INSTALL_PATH = /bin
TARGET_PATH = $(INSTALL_PATH)/$(APP_NAME)

USER = root
HOST = 192.168.1.32
EXPORT = export PORT=8080; export WEBSOCKET_PORT=8081

USER_HOST = "$(USER)@$(HOST)"
SSH_CMD = ssh $(USER_HOST)
CD_INSTALL_PATH = cd $(INSTALL_PATH)

deploy: deploy-stop deploy-clean build-jar deploy-upload deploy-start

deploy-upload:
	$(SSH_CMD) "mkdir -p $(INSTALL_PATH)"
	scp build/libs/$(APP_NAME_VERSION) "$(USER)@$(HOST):$(TARGET_PATH)"
	$(SSH_CMD) "chmod +x $(TARGET_PATH)"

deploy-start:
	$(SSH_CMD) "$(CD_INSTALL_PATH); $(EXPORT); setsid sh -c '/usr/bin/java -jar $(APP_NAME) > /dev/null 2>&1 &'"
	$(SSH_CMD) "ps ax | grep $(APP_NAME); $(CD_INSTALL_PATH); cat $(APP_NAME).pid"

deploy-stop:
	-$(SSH_CMD) "ps ax | grep $(APP_NAME) | grep -v grep | sed -e 's/\s.*//g' | xargs kill"

deploy-clean:
	-($(SSH_CMD) "rm $(TARGET_PATH)")
