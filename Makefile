GO          := go
GOBUILD     := $(GO) build
GOTEST      := $(GO) test
GOCLEAN     := $(GO) clean
GOINSTALL   := $(GO) install
BINARY_NAME := gex

.PHONY: all build install test clean

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) .

install: build
	$(GOINSTALL) .

test:
	$(GOTEST) ./...
	./test.sh

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
