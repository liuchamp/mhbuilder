.PHONY: build build-alpine clean test help default version

BIN_NAME=mhbuilder

OLD_VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
OLD_BUILD_DATE := $(shell grep "const BuildDate " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
OLD_COMMIT := $(shell grep "const GitCommit " version/version.go | sed -E 's/.*"(.+)"$$/\1/')

GIT_VERSION :=$(shell git describe --tags `git rev-list --tags --max-count=1`)
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "lacion/mhbuilder"

default: test

help:
	@echo 'Management commands for mhbuilder:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs dep ensure, mostly used for ci.'
	@@echo '   version update'
	@echo '    make clean           Clean the directory tree.'
	@echo



build: version
	@echo "building ${BIN_NAME} ${GIT_VERSION} ${BUILD_DATE} ${GIT_COMMIT}"
	@echo "GOPATH=${GOPATH}"
	@echo ""
	go build -ldflags "-X github.com/liuchamp/mhbuilder/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/liuchamp/mhbuilder/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

get-deps:
	dep ensure

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test -v -cover -coverprofile=coverage.data ./...
	go tool cover -html=coverage.data -o coverage.html &&  go tool cover -func=coverage.data -o coverage.txt

version:
	@echo 'update build info'
	@echo  '${OLD_VERSION} ${GIT_VERSION}'
	@echo  '${OLD_BUILD_DATE} ${BUILD_DATE}'
	@echo  '${OLD_COMMIT} ${GIT_COMMIT}'
	sed -i "" "s/${OLD_VERSION}/${GIT_VERSION}/g; s/${OLD_BUILD_DATE}/${BUILD_DATE}/g; s/${OLD_COMMIT}/${GIT_COMMIT}/g" version/version.go
