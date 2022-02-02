.PHONY: all test

all: build test
build:
		@go build

test:
		@go test

install: build
	@go install

release:
	@env GOOS=linux go build
	@env GOOS=darwin go build
