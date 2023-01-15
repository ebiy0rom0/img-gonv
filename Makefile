GOOS:=$(shell go env GOOS)
ifeq ($(GOOS), windows)
	SHELL:=powershell.exe
endif

GOCMD:=go
GOBUILD:=go build

CMD_DIR:=./cmd/gonv
CONF_NAME:=gonv.cnf

BIN_DIR:=./bin
BIN_NAME:=gonv
ifeq ($(GOOS), windows)
	BIN_NAME:=$(BIN_NAME).exe
endif
BIN_PATH:=$(BIN_DIR)/$(BIN_NAME)


PHONY: help
help:
	@echo help

PHONY: build
build: $(BIN_PATH) copy

$(BIN_PATH): mkdir clean
	$(GOBUILD) -o $@ -ldflags '-s -w' $(CMD_DIR)/main.go

PHONY: mkdir
mkdir:
ifeq ($(GOOS), windows)
	mkdir -Force $(BIN_DIR) > $$null
else
	mkdir -p $(BIN_DIR)
endif

PHONY: clean
clean:
ifeq ($(GOOS), windows)
	Remove-Item -Recurse -Force -Path $(BIN_DIR)/*
else
	rm -rf $(BIN_DIR)/*
endif

PHONY: copy
copy:
ifeq ($(GOOS), windows)
	Copy-Item $(CMD_DIR)/$(CONF_NAME) -Force -Destination $(BIN_DIR) > $$null
else
	cp -f $(CMD_DIR)/$(CONF_NAME) $(BIN_DIR)
endif
