.PHONY: release;
.ONESHELL: ;             # recipes execute in same shell

RELEASE_DIR = Release

APP_NAME = IDCardAndFingerprintReader

UNIFIED_NAME := $(APP_NAME)
ifeq ($(OS), Windows_NT)
UNIFIED_NAME := $(UNIFIED_NAME).exe
endif

all: release

release:
	-(rm -r $(RELEASE_DIR))
	-(mkdir $(RELEASE_DIR))
	cp build-$(APP_NAME)-Desktop_Qt_5_12_6_MinGW_32_bit-Release/release/$(UNIFIED_NAME) $(RELEASE_DIR)/$(UNIFIED_NAME)
	cp $(APP_NAME)/lib/* release/
	cd $(RELEASE_DIR)
	@echo windeployqt $(UNIFIED_NAME)
	cmd /A /Q /K C:/Developer/qt/5.12.6/mingw73_32/bin/qtenv2.bat
