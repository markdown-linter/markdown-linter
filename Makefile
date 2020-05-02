GO=go
GOTEST=$(GO) test
GOCOVER=$(GO) tool cover

all: build

clean:
	rm -rf ./bin/*

dependencies:
	$(GO) mod download

build: dependencies build-cmd

build-cmd:
	$(GO) build -o ./bin/markdown-linter ./main.go

test:
	$(GOTEST) -v -test.failfast ./...

ci: dependencies test

cover:
	$(GOTEST) -v -coverprofile=/tmp/coverage.out ./...
	$(GOCOVER) -func=/tmp/coverage.out
	$(GOCOVER) -html=/tmp/coverage.out
