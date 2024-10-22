SHELL :=/bin/bash

Version := $(shell git tag --sort=committerdate | tail -1)
LDFLAGS := "-s -w -X github.com/yankeexe/slack-status-cli/cmd.Version=$(Version)"

.DEFAULT_GOAL=help

test: # Run all tests
	@go test ./... -v

cover: # Launch a browser tab with coverage report
	@go test -v -coverprofile=c.out ./...
	@go tool cover -html=c.out
.PHONY: cover

clean: # Delete temp files
	@rm -f c.out || true
	@echo Temp files cleaned!

setup: # go mod tidy
	@go mod tidy
.PHONY: setup

build-amd64:
	@echo "Compiling amd64 build"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/st-amd64

build-armhf:
	@echo "Compiling armhf build"
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/st-armhf

build-arm64:
	@echo "Compiling arm64 build"
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/st-arm64

build-darwin:
	@echo "Compiling darwin build"
	@CGO_ENABLED=0 GOOS=darwin go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/st-darwin

build-darwin-arm:
	@echo "Compiling darwin ARM build"
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/st-darwin

build-windows:
	@echo "Compiling windows build"
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/st.exe

dist: # Create distribution binaries
	@mkdir -p bin
	@$(MAKE) -s build-amd64
	@$(MAKE) -s build-armhf
	@$(MAKE) -s build-arm64
	@$(MAKE) -s build-darwin
	@$(MAKE) -s build-darwin-arm
	@$(MAKE) -s build-windows


help: # Show this help
	@egrep -h '\s#\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
