HOMEDIR := $(shell pwd)
REPO := ""
DATEINFO_TAG ?= $(shell TZ='Asia/Shanghai' date +%Y%m%d%H%M%S)
BUILDINFO_TAG ?= $(shell git rev-parse --short HEAD)

GOCMD := go1.23
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
BINARY_NAME := http-echo

GO_BUILDINFO = -X 'http-echo/lib/buildinfo.Version=$(DATEINFO_TAG)-$(BUILDINFO_TAG)'

all: build

build:
	@echo "==> Building..."
	$(GOBUILD) -ldflags "$(GO_BUILDINFO)" -o $(BINARY_NAME) -v $(HOMEDIR)


test:
	$(GOCMD) test -race -timeout=120s -v -cover $(HOMEDIR)
