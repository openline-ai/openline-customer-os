run: build
	go run events-processing-platform.go

generate:

build: generate
	go build -v

.PHONY: test
test:
	go test -v ./... -cover

install:
	go install gotest.tools/gotestsum@latest

clean:

all: install build
