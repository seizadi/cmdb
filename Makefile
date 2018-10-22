PROJECT_ROOT		:= github.com/seizadi/cmdb
BUILD_PATH		:= bin
DOCKERFILE_PATH		:= $(CURDIR)/docker

# configuration for image names
USERNAME		:= $(USER)
GIT_COMMIT		:= $(shell git describe --dirty=-unsupported --always || echo pre-commit)
IMAGE_VERSION		?= $(USERNAME)-dev-$(GIT_COMMIT)
IMAGE_REGISTRY ?= soheileizadi/cmdb

# configuration for server binary and image
SERVER_BINARY 		:= $(BUILD_PATH)/server
SERVER_PATH		:= $(PROJECT_ROOT)/cmd/server
SERVER_IMAGE		:= $(IMAGE_REGISTRY)/cmdb-server
SERVER_DOCKERFILE	:= $(DOCKERFILE_PATH)/Dockerfile

# configuration for the protobuf gentool
SRCROOT_ON_HOST		:= $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
SRCROOT_IN_CONTAINER	:= /go/src/$(PROJECT_ROOT)
DOCKER_RUNNER    	:= docker run -u `id -u`:`id -g` --rm
DOCKER_RUNNER		+= -v $(SRCROOT_ON_HOST):$(SRCROOT_IN_CONTAINER)
DOCKER_GENERATOR	:= infoblox/atlas-gentool:v3
GENERATOR		:= $(DOCKER_RUNNER) $(DOCKER_GENERATOR)

# configuration for the database
DATABASE_ADDRESS	?= localhost:5432

# configuration for building on host machine
GO_CACHE		:= -pkgdir $(BUILD_PATH)/go-cache
GO_BUILD_FLAGS		?= $(GO_CACHE) -i -v
GO_TEST_FLAGS		?= -v -cover
GO_PACKAGES		:= $(shell go list ./... | grep -v vendor)

.PHONY: all
all: vendor protobuf docker

.PHONY: fmt
fmt:
	@go fmt $(GO_PACKAGES)

.PHONY: test
test: fmt
	@go test $(GO_TEST_FLAGS) $(GO_PACKAGES)

.PHONY: docker
docker:
	@docker build -f $(SERVER_DOCKERFILE) -t $(SERVER_IMAGE):$(IMAGE_VERSION) .
	@docker image prune -f --filter label=stage=server-intermediate
.PHONY: push
push:
	@docker push $(SERVER_IMAGE)

.PHONY: protobuf
protobuf:
	@$(GENERATOR) \
	--go_out=plugins=grpc:. \
	--grpc-gateway_out=logtostderr=true:. \
	--gorm_out=. \
	--validate_out="lang=go:." \
	--swagger_out=:. $(PROJECT_ROOT)/pkg/pb/service.proto

.PHONY: vendor
vendor:
	@dep ensure -vendor-only

.PHONY: vendor-update
vendor-update:
	@dep ensure
	
.PHONY: clean
clean:
	@docker rmi -f $(shell docker images -q $(SERVER_IMAGE)) || true

.PHONY: migrate-up
migrate-up:
	@migrate -database 'postgres://$(DATABASE_ADDRESS)/cmdb?sslmode=disable' -path ./db/migrations up
	
.PHONY: migrate-down
migrate-down:
	@migrate -database 'postgres://$(DATABASE_ADDRESS):5432/cmdb?sslmode=disable' -path ./db/migrations down

