#@IgnoreInspection BashAddShebang

ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
BUILD_INFO_PKG="github.com/thanos-go/build_info"
BUILD_OUTPUT ?= ./build/thanos
BUILD_MAC_OUTPUT ?= ./build/thanos_mac
CI_COMMIT_REF_NAME ?= $(shell git symbolic-ref -q --short HEAD || git describe --tags --exact-match)
CI_COMMIT_SHORT_SHA ?= $(shell git rev-parse --verify --short=8 HEAD)
CURRENT_DATETIME=$(shell TZ=Asia/Tehran date '+%FT%T')
LDFLAGS="-w -s -X $(BUILD_INFO_PKG).BuildTime=$$(TZ=Asia/Tehran date '+%FT%T') -X $(BUILD_INFO_PKG).AppVersion=$$(git rev-parse HEAD | cut -c 1-8) -X $(BUILD_INFO_PKG).VCSRef=$$(git rev-parse --abbrev-ref HEAD)"

.PHONY: all build

fix: lint format
all: fix build

############################################################
## Compile and Build
############################################################


build:
	go build -o $(BUILD_OUTPUT) -ldflags $(LDFLAGS) -tags=jsoniter .

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o $(BUILD_OUTPUT) -ldflags $(LDFLAGS) .

compile-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod=vendor -o $(BUILD_MAC_OUTPUT) -ldflags $(LDFLAGS) .

docker-build:
	docker build --build-arg "BUILD_DATE=$(CURRENT_DATETIME)" --build-arg "VCS_REF=$(CI_COMMIT_SHORT_SHA)" --build-arg "BUILD_VERSION=$(CI_COMMIT_REF_NAME)" -t thanos:latest .

run:
	go run --ldflags $(LDFLAGS) . serve

fresh:
	go run --ldflags $(LDFLAGS) . migrate fresh

############################################################
## Format and Lint
############################################################

format:
	which goimports || GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R goimports -w R
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R gofmt -s -w R

lint:
	which golangci-lint || (GO111MODULE=off go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint)
	golangci-lint run --skip-files ".*_test.go" --deadline 10m $(ROOT)/...

############################################################
## Test
############################################################

test:
	go test ./...

ci-test:
	go test -coverprofile=build/coverage.txt ./...
	go tool cover -func build/coverage.txt

############################################################
## Dev
############################################################

dev:
	which CompileDaemon || (GO111MODULE=off go get github.com/githubnemo/CompileDaemon)
	HOST_IP=127.0.0.1 CompileDaemon -color=true -log-prefix=false --build="make build" -exclude-dir="build/mysql" -exclude-dir="vendor" --command="$(BUILD_OUTPUT) serve"

up:
	docker-compose up -d
	make dev

down:
	docker-compose down