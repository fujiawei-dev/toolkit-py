{{MAKEFILE_HEADER}}

.PHONY: ;
#.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.IGNORE: clean;            # ignore all errors, keep going

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

BUILD_CMD := go build -trimpath -ldflags $(BUILD_ARGS)

DEBUG_CMD := ./$(UNIFIED_NAME) -D -S -c storage/configs/settings-debug.yaml -d
RELEASE_CMD := ./$(UNIFIED_NAME) -D -S -P -c storage/configs/settings-release.yaml -d

all: tidy swag start

tidy:
	go mod tidy

swag:
	go get -u github.com/swaggo/swag
	-(export PATH=$(go env GOPATH)/bin:$PATH)
	-(swag init)

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
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

darwin-arm64:
	GOARCH=arm64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-386:
	GOARCH=386 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-armv5:
	GOARCH=arm GOOS=linux GOARM=5 $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-armv6:
	GOARCH=arm GOOS=linux GOARM=6 $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-armv7:
	GOARCH=arm GOOS=linux GOARM=7 $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-armv8:
	GOARCH=arm64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-mips-softfloat:
	GOARCH=mips GOMIPS=softfloat GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-mips-hardfloat:
	GOARCH=mips GOMIPS=hardfloat GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-mipsle-softfloat:
	GOARCH=mipsle GOMIPS=softfloat GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-mipsle-hardfloat:
	GOARCH=mipsle GOMIPS=hardfloat GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-mips64:
	GOARCH=mips64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

linux-mips64le:
	GOARCH=mips64le GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

freebsd-386:
	GOARCH=386 GOOS=freebsd $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

freebsd-amd64:
	GOARCH=amd64 GOOS=freebsd $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

freebsd-arm64:
	GOARCH=arm64 GOOS=freebsd $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

windows-386:
	GOARCH=386 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

windows-arm32v7:
	GOARCH=arm GOOS=windows GOARM=7 $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

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
	rm $(BINDIR)/*
