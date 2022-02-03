SHELL = /bin/bash

all: build test

build:
	go build ./...

test:
	go test -race ./.

e2e:
	go run ./test/e2e

gen:
	GO_POST_PROCESS_FILE="gofmt -w" \
		openapi-generator batch --clean --verbose --includes-base-dir apis/ \
		apis/private/v1alpha1/gen.yaml \
		apis/public/v1/gen.yaml
	go mod tidy

.PHONY: all build test e2e gen
