# Makefile for OCPP Example project
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD)fmt

# Folders definitions
BINARY_FOLDER=bin

# Get a short hash of the git had for building images.
VERSION = 1.0
TAG = $$(git rev-parse --short HEAD)
APP_NAME=golib
IMAGE_NAME = "${APP_NAME}"
DOCKERFILE = "Dockerfile"
PROJECT_PATH = $(shell pwd)

define docker-build =
	docker build -t ${IMAGE_NAME}:${VERSION} -f ${DOCKERFILE} .
	docker run --rm \
		--name toolbox-golib-tests \
		-v ${PROJECT_PATH}:/go/golib \
		${IMAGE_NAME}:${VERSION} \
		make -C /go/golib $(1)
endef

define docker-run =
	docker run --rm \
		--name toolbox-golib-tests \
		-v ${PROJECT_PATH}:/go/golib \
		-v ${PROJECT_PATH}/logs:/tmp/logs \
		-p "9033:8080" \
		${IMAGE_NAME}:${VERSION} \
		$(1)
endef

fmts:
	$(GOFMT) -s -d ./timelib/*.go
	$(GOFMT) -s -d ./logging/*.go
	$(GOFMT) -s -d ./tools/*.go

test:
	@CGO_ENABLED=0 $(GOTEST) -v ./...

modinit:
	$(GOMOD) init github.com/CoderSergiy/golib

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_FOLDER)

testlib: fmts test

# Command using Docker container

dockerclean:
	$(call docker-build, "clean")

dockertest:
	$(call docker-build, "testlib")

