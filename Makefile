GO=go
GOTEST=$(GO) test
GOCOVER=$(GO) tool cover

all: build

clean:
	rm -rf ./bin/*

dependencies:
	$(GO) mod download

build: clean dependencies build-cmd

build-cmd:
	sh -c "'$(CURDIR)/scripts/build.sh'"

test:
	$(GOTEST) -v -test.failfast ./...

ci: dependencies test

cover:
	$(GOTEST) -v -coverprofile=/tmp/coverage.out ./...
	$(GOCOVER) -func=/tmp/coverage.out
	$(GOCOVER) -html=/tmp/coverage.out
