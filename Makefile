SHELL = /bin/bash

all: build test

build:
	go build ./...

test:
	go test -race ./.

gen:
	GO_POST_PROCESS_FILE="gofmt -w" \
		openapi-generator batch --clean --verbose --includes-base-dir apis/ \
		apis/public/v1/gen.yaml
	go mod tidy

.PHONY: build test gen
