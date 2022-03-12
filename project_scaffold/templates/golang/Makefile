{{HASHTAG_COMMENTS}}

.PHONY: ;
#.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.IGNORE: clean swag;            # ignore all errors, keep going

APP_VERSION = 0.0.1
APP_NAME = {{APP_NAME}}
APP_MODULE = {{GOLANG_MODULE}}

UNIFIED_NAME := $(APP_NAME)
ifeq ($(OS), Windows_NT)
UNIFIED_NAME := $(UNIFIED_NAME).exe
endif

GIT_COMMIT := $(shell git rev-parse --short HEAD || echo unsupported)
BUILD_TIME := $(shell date --rfc-3339 seconds  | sed -e 's/ /T/' || echo unsupported)

# go tool link --help
BUILD_ARGS := "\
-X $(APP_MODULE)/version.GitCommit=$(GIT_COMMIT) \
-X $(APP_MODULE)/version.Name=$(APP_NAME) \
-X $(APP_MODULE)/version.Version=$(APP_VERSION) \
-X $(APP_MODULE)/version.BuildTime=$(BUILD_TIME) \
-w -s -buildid=\
"
# The -w and -s flags reduce binary sizes by excluding unnecessary symbols and debug info
# The -buildid= flag makes builds reproducible

# CGO cross is not supported (CGO_ENABLED=1)
BUILD_CMD := go build -trimpath -ldflags $(BUILD_ARGS)

DEBUG_CMD := ./$(UNIFIED_NAME) -D -S -c storage/configs/settings-debug.yaml -d
RELEASE_CMD := ./$(UNIFIED_NAME) -D -S -P -c storage/configs/settings-release.yaml -d

all: tidy swag start

tidy dep:
	go mod tidy

init:
	git init
	git add .
	git commit -m "first commit"

swag:
	go get github.com/swaggo/swag
	-(swag init; swag fmt)

build:
	$(BUILD_CMD) -o $(UNIFIED_NAME) main.go

start: build
	$(DEBUG_CMD) start

stop:
	$(DEBUG_CMD) stop

config:
	$(DEBUG_CMD) config

restart: stop start

#------------------------------------------------------------
# Release

BIN_DIR = bin

builds: linux-amd64 darwin-amd64 windows-amd64 # Most used

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

darwin-arm64:
	GOARCH=arm64 GOOS=darwin $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-386:
	GOARCH=386 GOOS=linux $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-amd64:
	GOARCH=amd64 GOOS=linux $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-armv5:
	GOARCH=arm GOOS=linux GOARM=5 $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-armv6:
	GOARCH=arm GOOS=linux GOARM=6 $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-armv7:
	GOARCH=arm GOOS=linux GOARM=7 $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-armv8:
	GOARCH=arm64 GOOS=linux $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-mips-softfloat:
	GOARCH=mips GOMIPS=softfloat GOOS=linux $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-mips-hardfloat:
	GOARCH=mips GOMIPS=hardfloat GOOS=linux $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-mipsle-softfloat:
	GOARCH=mipsle GOMIPS=softfloat GOOS=linux $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-mipsle-hardfloat:
	GOARCH=mipsle GOMIPS=hardfloat GOOS=linux $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-mips64:
	GOARCH=mips64 GOOS=linux $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

linux-mips64le:
	GOARCH=mips64le GOOS=linux $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

freebsd-386:
	GOARCH=386 GOOS=freebsd $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

freebsd-amd64:
	GOARCH=amd64 GOOS=freebsd $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

freebsd-arm64:
	GOARCH=arm64 GOOS=freebsd $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@

windows-386:
	GOARCH=386 GOOS=windows $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@.exe

windows-amd64:
	GOARCH=amd64 GOOS=windows $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@.exe

windows-arm32v7:
	GOARCH=arm GOOS=windows GOARM=7 $(BUILD_CMD) -o $(BIN_DIR)/$(APP_NAME)-$@.exe

PLATFORM_LIST = \
	darwin-amd64 \
	darwin-arm64 \
	linux-386 \
	linux-amd64 \
	linux-armv5 \
	linux-armv6 \
	linux-armv7 \
	linux-armv8 \
	linux-mips-softfloat \
	linux-mips-hardfloat \
	linux-mipsle-softfloat \
	linux-mipsle-hardfloat \
	linux-mips64 \
	linux-mips64le \
	freebsd-386 \
	freebsd-amd64 \
	freebsd-arm64

WINDOWS_ARCH_LIST = \
	windows-386 \
	windows-amd64 \
	windows-arm32v7

all-arch: $(PLATFORM_LIST) $(WINDOWS_ARCH_LIST)

clean:
	rm $(BIN_DIR)/*

tag:
	git tag v$(APP_VERSION)
	git push origin v$(APP_VERSION)
