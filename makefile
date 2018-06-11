# Project Name
SHA1         		:= $(shell git rev-parse --verify --short HEAD)
MAJOR_VERSION		:= $(shell cat package.json | sed -n 's/.*"major": "\(.*\)",/\1/p')
MINOR_VERSION		:= $(shell cat package.json | sed -n 's/.*"minor": "\(.*\)"/\1/p')
INTERNAL_BUILD_ID	:= $(shell [ -z "${BUILD_ID}" ] && echo "0" || echo ${BUILD_ID})
PROJECT				:= $(shell cat package.json | sed -n 's/.*"name": "\(.*\)",/\1/p')
VERSION				:= $(shell echo "${MAJOR_VERSION}-${MINOR_VERSION}-${INTERNAL_BUILD_ID}-${SHA1}")
DOCKER_VERSION		:= $(shell echo "${MAJOR_VERSION}.${MINOR_VERSION}.${INTERNAL_BUILD_ID}-${SHA1}")
BUILD_NAME			:= $(shell echo "go-build")
BUILD_IMAGE			:= $(shell echo "golang:1.10")
PWD					:= $(shell pwd)
REPO 				:= $(shell echo "github.com/cameronnewman")

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := build

.PHONY: tests
tests:
	docker run -t -i --rm --name=$(BUILD_NAME) \
	-t -v $(PWD):/go/src/$(REPO)/$(PROJECT) \
	-w /go/src/$(REPO)/$(PROJECT) $(BUILD_IMAGE) \
	go test -v ./...

.PHONY: build
build: tests
	docker run -t -i --rm \
	--name=$(BUILD_NAME) \
	-e "CGO_ENABLED=0" \
	-e "GOOS=linux" -t \
	-v $(PWD):/go/src/$(REPO)/$(PROJECT) \
	-w /go/src/$(REPO)/$(PROJECT) $(BUILD_IMAGE) \
	go build -x -ldflags "-X main.version=$(VERSION)" \
	-a -installsuffix cgo -o $(PROJECT) iex.go

