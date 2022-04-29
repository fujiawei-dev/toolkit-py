APP_NAME = $(shell ls *.iml | sed -e 's/\.iml/.jar/g' | sed -e 's/-.*-all//g')

echo:
	@echo $(shell cd build/libs; ls *.jar)
	@echo $(APP_NAME)

build-jar:
	-rm -rf build/libs
	./gradlew.bat shadowJar #--status

INSTALL_PATH = /bin
TARGET_PATH = $(INSTALL_PATH)/$(APP_NAME)

USER = root
HOST = 192.168.1.32
EXPORT = export PORT=8080; export WEBSOCKET_PORT=8081

USER_HOST = "$(USER)@$(HOST)"
SSH_CMD = ssh $(USER_HOST)
CD_INSTALL_PATH = cd $(INSTALL_PATH)

deploy: deploy-stop deploy-clean deploy-upload deploy-start

deploy-upload: build-jar
	$(SSH_CMD) "mkdir -p $(INSTALL_PATH)"
	scp $(shell ls build/libs/*.jar) "$(USER)@$(HOST):$(TARGET_PATH)"
	$(SSH_CMD) "chmod +x $(TARGET_PATH)"

deploy-start:
	$(SSH_CMD) "$(CD_INSTALL_PATH); $(EXPORT); setsid sh -c '/usr/bin/java -jar $(APP_NAME) > /dev/null 2>&1 &'"
	$(SSH_CMD) "ps ax | grep $(APP_NAME)"

deploy-stop:
	-$(SSH_CMD) "ps ax | grep $(APP_NAME) | grep -v grep | sed 's/^\s*//g' | sed -e 's/\s.*//g' | xargs kill"

deploy-clean:
	-($(SSH_CMD) "rm $(TARGET_PATH)")
