.PHONY: all test

all: build test
build:
		@gb build

test:
		@gb test

install: build
	@cp bin/csvq $(GOPATH)/bin/csvq

release:
	@env GOOS=linux gb build
	@env GOOS=darwin gb build
