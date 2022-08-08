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
-X $(APP_MODULE)/version.GitCommit=$(GIT_COMMIT) \
-X $(APP_MODULE)/version.Name=$(APP_NAME) \
-X $(APP_MODULE)/version.Version=$(APP_VERSION) \
-X $(APP_MODULE)/version.BuildTime=$(APP_BUILD_TIME) \
-X $(APP_MODULE)/version.MakeVersion=$(MAKE_VERSION) \
-w -s -buildid=\
"

AIR_PACKAGE ?= github.com/cosmtrek/air@latest

SWAG_PACKAGE ?= github.com/swaggo/swag/cmd/swag@latest

# CGO cross is not supported (CGO_ENABLED=1)
GO_BUILD := go build -trimpath -ldflags $(LDFLAGS)

RUN_ARGS := -D -S -P -c config/settings-debug.yaml

BIN_DIR = bin

all: build

tidy:
	go mod tidy
	go install $(AIR_PACKAGE)
	go install $(SWAG_PACKAGE)

swag:
	swag init; swag fmt

build: tidy swag go-build

run: swag
	go run main.go $(RUN_ARGS) start

watch: swag
	go run $(AIR_PACKAGE) -c .air.toml $(RUN_ARGS) start

tag:
	git tag v$(APP_VERSION)
	git push origin v$(APP_VERSION)

go-build:
	$(GO_BUILD) -o $(EXECUTABLE) main.go

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-armv8:
	GOARCH=arm64 GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME)-$@

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GO_BUILD) -o $(BIN_DIR)/$(APP_NAME)-$@.exe
