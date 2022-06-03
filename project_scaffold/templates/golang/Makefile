{{HASHTAG_COMMENTS}}

.PHONY: ;
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell


APP_NAME = gin_template
APP_VERSION = 0.0.1

GO_MODULE = gin-template

ifeq ($(OS), Windows_NT)
	SHELL = pwsh
	SHELLFLAGS = -Command
	EXECUTABLE ?= ${APP_NAME}.exe
else
	SHELL = bash
	SHELLFLAGS = -c
	EXECUTABLE ?= ${APP_NAME}
endif

GIT_COMMIT := $(shell git rev-parse --short HEAD || echo unsupported)
BUILD_TIME := $(shell date --rfc-3339 seconds  | sed -e 's/ /T/' || echo unsupported)

# go tool link --help
# The -w and -s flags reduce binary sizes by excluding unnecessary symbols and debug info
# The -buildid= flag makes builds reproducible
LDFLAGS := "\
-X $(GO_MODULE)/version.GitCommit=$(GIT_COMMIT) \
-X $(GO_MODULE)/version.Name=$(APP_NAME) \
-X $(GO_MODULE)/version.Version=$(APP_VERSION) \
-X $(GO_MODULE)/version.BuildTime=$(BUILD_TIME) \
-w -s -buildid=\
"

# CGO cross is not supported (CGO_ENABLED=1)
GO_BUILD := go build -trimpath -ldflags $(LDFLAGS)
RUN_ARGS :=  -D -S -P -c storage/configs/settings-debug.yaml

BIN_DIR = bin
AIR_PACKAGE ?= github.com/cosmtrek/air@latest
SWAG_PACKAGE ?= github.com/swaggo/swag/cmd/swag@latest

all: build

.PHONY: help
help:
	@echo "Make Routines:"
	@echo " - \"\"                               equivalent to \"build\""
	@echo " - build                            build everything"
	@echo " - backend                          build backend files"
	@echo " - watch                            watch everything and continuously rebuild"
	@echo " - watch-backend                    watch backend files and continuously rebuild"
	@echo " - clean                            delete backend and integration files"
	@echo " - deps                             install dependencies"
	@echo " - deps-backend                     install backend dependencies"
	@echo " - test                             test everything"
	@echo " - test-backend                     test backend files"
	@echo " - generate-swagger                 generate the swagger spec from code comments"

.PHONY: build
build: backend

.PHONY: deps
deps: deps-backend

deps-backend:
	go mod download
	go install $(AIR_PACKAGE)
	go install $(SWAG_PACKAGE)

go-tidy:
	go mod tidy
	@diff=$$(git diff go.sum); \
	if [ -n "$$diff" ]; then \
		echo "Please run '$(GO) mod tidy' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi

go-build:
	$(GO_BUILD) -o $(EXECUTABLE) main.go

.IGNORE: generate-swagger
generate-swagger:
	swag init; swag fmt

.PHONY: backend
backend: go-tidy generate-swagger go-build

watch:
	watch-backend

watch-backend: generate-swagger
	# air -c .air.toml
	go run $(AIR_PACKAGE) -c .air.toml $(RUN_ARGS)

go-run:
	go run main.go $(RUN_ARGS) start

builds: linux-amd64 darwin-amd64 windows-amd64 # Most used

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-armv8:
	GOARCH=arm64 GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME)-$@

windows-386:
	GOARCH=386 GOOS=windows $(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME)-$@.exe

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME)-$@.exe

git-tag:
	git tag v$(APP_VERSION)
	git push origin v$(APP_VERSION)

git-init:
	-(rm internal/.gitignore)
	git init; git add .; git commit -m "first commit"

# Remote Deployment
USER = root
HOST = raspberrypi
USER_HOST = "$(USER)@$(HOST)"
SSH_REMOTE = ssh $(USER_HOST)

INSTALL_PATH = ~/bin
TARGET_ARCH = linux-amd64
TARGET_PATH = $(INSTALL_PATH)/$(APP_NAME)
CONFIG_FILE = storage/configs/deploy/debug.yaml
CD_INSTALL_PATH = cd $(INSTALL_PATH)
DEPLOY_ARGS :=  -D -S -P -d -c settings.yaml

# make -n deploy
deploy: deploy-stop deploy-clean $(TARGET_ARCH) deploy-upload deploy-start

deploy-upload:
	$(SSH_REMOTE) "mkdir -p $(INSTALL_PATH)"
	scp $(BIN_DIR)/$(APP_NAME)-$(TARGET_ARCH) "$(USER)@$(HOST):$(TARGET_PATH)"
	scp $(CONFIG_FILE) "$(USER)@$(HOST):$(INSTALL_PATH)/settings.yaml"
	$(SSH_REMOTE) "mkdir -p $(INSTALL_PATH)/assets"
	scp -r assets "$(USER)@$(HOST):$(INSTALL_PATH)"
	$(SSH_REMOTE) "chmod +x $(TARGET_PATH)"

deploy-clean:
	-($(SSH_REMOTE) "rm $(TARGET_PATH)")

deploy-start:
	$(SSH_REMOTE) "$(CD_INSTALL_PATH); $(TARGET_PATH) $(DEPLOY_PATH) start"
	$(SSH_REMOTE) "ps -ef | grep $(APP_NAME)"

deploy-stop:
	-($(SSH_REMOTE) "$(CD_INSTALL_PATH); $(TARGET_PATH) stop")

deploy-reset: deploy-stop
	$(SSH_REMOTE) "$(CD_INSTALL_PATH); $(TARGET_PATH) $(DEPLOY_PATH) reset"
