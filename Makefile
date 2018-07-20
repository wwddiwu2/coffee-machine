.PHONY: build test

all: test build

build:
	go install ./...

test:
	go test -v ./...