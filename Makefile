.PHONY: install build

install:
	@go mod download

build:
	@go build .

build-mac:
	GOOS=darwin GOARCH=amd64 go build .
