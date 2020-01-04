VERSION := $(shell git describe --always --long --dirty)

buildrelease: build compress
.PHONY: buildrelease

build: binary
.PHONY: build

binary:
	echo "building binary"
	go build -ldflags="-s -w -X main.buildVersion=${VERSION}" -o BackUploader

compress:
	echo "compressing binary"
	upx -9 --brute output/linux-amd64/FonFon
