run: build
	go run integrity_checker_main.go

generate:
	export PATH="${PATH}:$(shell go env GOPATH)/bin";

build: generate
	go build -v -o bin/integrity-checker

.PHONY: test
test:
	go test -v ./...

install:
	go install gotest.tools/gotestsum@latest

clean:
	find gen -mindepth 1 -maxdepth 1 -exec rm -r {} \;
	rm -rf bin

all: install build
