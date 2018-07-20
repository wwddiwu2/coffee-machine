.PHONY: build test

all: build test

build:
	vgo install ./...

test:
	vgo test -v ./...