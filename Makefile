.PHONY: build test

all: test build

build:
	go install ./...

test:
	go get -t ./...
	go test -v ./...