VERSION := $(shell git describe --always --long --dirty)

buildrelease: build compress
.PHONY: buildrelease

build: binary
.PHONY: build

jenkins:
	go build -ldflags="-s -w -X main.buildVersion=${VERSION} -X main.buildArch=${GOOS}-${GOARCH}" -o BackUploader.${GOOS}.${GOARCH}

prepare:
	go get -v golang.org/x/sync/semaphore
	go get -v github.com/wsxiaoys/terminal
	go get -v github.com/bigkevmcd/go-configparser

binary:
	go build -ldflags="-s -w -X main.buildVersion=${VERSION}" -o BackUploader

compress:
	upx -9 --brute output/linux-amd64/FonFon
