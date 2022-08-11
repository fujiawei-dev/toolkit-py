.PHONY: build install;

# .ONESHELL: ; # recipes execute in same shell
# .NOTPARALLEL: ; # wait for target to finish
# .EXPORT_ALL_VARIABLES: ; # send all vars to shell


APP_NAME = {{ project_slug.snake_case }}

{%- if not open_source %}
APP_MODULE = {{ project_slug.kebab_case }}
{%- else %}
APP_MODULE = https://github.com/{{ github_username }}/{{ project_slug.kebab_case }}
{%- endif %}

APP_VERSION = 0.1.0

APP_BUILD_TIME := $(shell date --rfc-3339 seconds | sed -e 's/ /T/' || echo unsupported)
# APP_BUILD_TIME := $(shell date -D "YYYY-MM-DD hh:mm:ss" | sed -e 's/ /T/' || echo unsupported)

GIT_COMMIT := $(shell git rev-parse --short HEAD || echo unsupported)

MAKE_VERSION := $(shell "$(MAKE)" -v | head -n 1)

ifeq ($(OS), Windows_NT)
	SHELL = pwsh
	SHELLFLAGS = -Command
	EXECUTABLE ?= ${APP_NAME}.exe
	INSTALL_DIR = C:/Developer/bin
else
	SHELL = bash
	SHELLFLAGS = -c
	EXECUTABLE ?= ${APP_NAME}
	INSTALL_DIR = /usr/local/bin
endif

# go tool link --help
# The -w and -s flags reduce binary sizes by excluding unnecessary symbols and debug info
# The -buildid= flag makes builds reproducible
LDFLAGS := "\
-X main.GitCommit=$(GIT_COMMIT) \
-X main.Name=$(APP_NAME) \
-X main.Version=$(APP_VERSION) \
-X main.BuildTime=$(APP_BUILD_TIME) \
-X main.MakeVersion=$(MAKE_VERSION) \
-w -s -buildid=\
"

# CGO cross is not supported (CGO_ENABLED=1)
GO_BUILD := go build -trimpath -ldflags $(LDFLAGS)

BIN_DIR = bin

all: build

tag:
	git tag v$(APP_VERSION)
	git push origin v$(APP_VERSION)

tidy:
	go mod tidy

build: tidy go-build

go-build:
	$(GO_BUILD) -o $(EXECUTABLE) main.go

install: build
	cp $(BUILD_DIR)/$(EXECUTABLE) $(INSTALL_DIR)/$(EXECUTABLE)
	$(EXECUTABLE) -v

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-armv8:
	GOARCH=arm64 GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME)-$@

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME)-$@.exe
