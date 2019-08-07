include .env
export

MAIN_VERSION:=$(shell git describe --always --abbrev=7 --tags || echo "0.1")
AUTHOR:=$(shell git --no-pager show -s --format='%an' ${MAIN_VERSION})
PLATFORM:=$(shell go env GOOS)
ARCH:=$(shell go env GOARCH)
GOPATH:=$(shell go env GOPATH)
GOBIN:=$(GOPATH)/bin
LDFLAGS:="-X github.com/jorgechato/api.jorgechato.com/utils.VERSION=${MAIN_VERSION} -X github.com/jorgechato/api.jorgechato.com/utils.AUTHOR=${AUTHOR}"
M = $(shell printf "\033[34;1m▶\033[0m")


.PHONY: build clean help default server vendor release test

default: help

vendor: ; $(info $(M) Download dev dependencies...)
	go get -u github.com/tebeka/go2xunit

build: ; $(info $(M) Building project...) ## Build server binary
	go build -ldflags ${LDFLAGS} -a -o api.jorgechato.com server.go

run: ; $(info $(M) Run project... ) ## Run server without building it
	go run -ldflags ${LDFLAGS} server.go

clean: ; $(info $(M) Removing generated files... ) ## Clean schema bind data
	rm -rf *.out *.xml release *.output

test: vendor ; $(info $(M) Run tests...) ## Start unit tests
	go test -race -cover -v -coverprofile=coverage.out ./... > test.output
	cat test.output | go2xunit -output tests.xml

server: ; $(info $(M) Starting development server...) ## Start development server
	realize start || echo "please install realize (go get github.com/oxequa/realize)"

unit: vendor ; $(info $(M) Run unit tests...) ## Start unit tests
	go list $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=60s -parallel=4

fmtcheck:
	@sh -c "'$(CURDIR)/lazy/gofmtcheck.sh'"

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
