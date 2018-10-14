SHELL:="/bin/bash"
MAIN_VERSION:=$(shell git describe --always --abbrev=7 --tags || echo "0.1")
AUTHOR:=$(shell git --no-pager show -s --format='%an' ${MAIN_VERSION})
PLATFORM:=$(shell go env GOOS)
ARCH:=$(shell go env GOARCH)
GOPATH:=$(shell go env GOPATH)
GOBIN:=$(GOPATH)/bin
LDFLAGS:="-X github.com/jorgechato/2019.jorgechato.go/utils.VERSION=${MAIN_VERSION} -X github.com/jorgechato/2019.jorgechato.go/utils.AUTHOR=${AUTHOR}"
M = $(shell printf "\033[34;1m▶\033[0m")


.PHONY: build clean dep schema help default server setup vendor

default: help

vendor: ; $(info $(M) Download dev dependencies...)
	go get github.com/oxequa/realize

setup: ; $(info $(M) Fetching github.com/golang/dep...)
	go get github.com/golang/dep/cmd/dep

build: dep ; $(info $(M) Building project...) ## Build server binary
	go build -ldflags ${LDFLAGS} -a -o server server.go	

clean: ; $(info $(M) [TODO] Removing generated files... ) ## Clean schema bind data
	$(RM) schema/bindata.go

dep: setup ; $(info $(M) Ensuring vendored dependencies are up-to-date...) ## Download dependencies
	dep ensure -v -vendor-only

schema: dep ; $(info $(M) Embedding schema files into binary...) ## Build schema
	go generate ./schema

server: vendor schema ; $(info $(M) Starting development server...) ## Start development server
	realize start
	# go run -ldflags ${LDFLAGS} server.go

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
