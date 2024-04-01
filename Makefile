.PHONY: setup fmt lint test unittest build install clean vulncheck
.DEFAULT_GOAL := build

setup:
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin

fmt:
	go fmt ./...
	goimports -w $(shell find . -type f -name "*.go")

lint:
	golangci-lint run

test:
	go test ./...

unittest:
	go test $(shell go list ./... | grep -x -v -e "github.com/shibataka000/kms/crypto" -e "github.com/shibataka000/kms/crypto/kms")

build:
	go build

install:
	go install

clean:
	go clean -testcache

vulncheck:
	govulncheck ./...
