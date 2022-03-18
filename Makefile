VERSION := $(shell git describe --always --long --dirty)
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
GOARCH := $(shell go env GOEXE)

build: binary
.PHONY: build

binary:
	go build -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.buildArch=${GOOS}-${GOARCH}"
