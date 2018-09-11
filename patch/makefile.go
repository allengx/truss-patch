package patch

var Makefile = `
# Makefile for Truss.
#
SHA := $(shell git rev-parse --short=10 HEAD)

MAKEFILE_PATH := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

# Build native Truss by default.
default: truss

dependencies:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/jteeuwen/go-bindata/...

# Generate go files containing the all template files in []byte form
gobindata:
	go generate github.com/tuneinc/truss/gengokit/template

# Install truss
truss: gobindata
	go install -ldflags '-X "main.Version=$(SHA)" -X "main.VersionDate=Mon May 28 22:12:59 UTC 2018"' github.com/tuneinc/truss/cmd/truss

# Run the go tests and the truss integration tests
test: test-go test-integration

test-go:
	go test -v ./...

test-integration:
	$(MAKE) -C cmd/_integration-tests

# Removes generated code from tests
testclean:
	$(MAKE) -C cmd/_integration-tests clean

.PHONY: testclean test-integration test-go test truss gobindata dependencies
    `
