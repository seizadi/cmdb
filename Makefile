PROJECT_ROOT    := github.com/seizadi/cmdb
BUILD_PATH      := bin
DOCKERFILE_PATH := $(CURDIR)/docker

# configuration for image names
USERNAME       := $(USER)
GIT_COMMIT     := $(shell git describe --dirty=-unsupported --always || echo pre-commit)
IMAGE_VERSION  ?= $(USERNAME)-dev-$(GIT_COMMIT)
IMAGE_REGISTRY ?= soheileizadi

# configuration for server binary and image
SERVER_BINARY     := $(BUILD_PATH)/server
SERVER_PATH       := $(PROJECT_ROOT)/cmd/server
SERVER_IMAGE      := $(IMAGE_REGISTRY)/cmdb-server
SERVER_DOCKERFILE := $(DOCKERFILE_PATH)/Dockerfile
# Placeholder. modify as defined conventions.
DB_VERSION        := 3
SRV_VERSION       := $(shell git describe --tags)
API_VERSION       := v1

# configuration for the protobuf gentool
SRCROOT_ON_HOST      := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
SRCROOT_IN_CONTAINER := /go/src/$(PROJECT_ROOT)
DOCKER_RUNNER        := docker run --rm
DOCKER_RUNNER        += -v $(SRCROOT_ON_HOST):$(SRCROOT_IN_CONTAINER)
DOCKER_GENERATOR     := infoblox/atlas-gentool:latest
GENERATOR            := $(DOCKER_RUNNER) $(DOCKER_GENERATOR)

# configuration for the database
DATABASE_HOST ?= localhost:5432

# configuration for building on host machine
GO_CACHE       := -pkgdir $(BUILD_PATH)/go-cache
GO_BUILD_FLAGS ?= $(GO_CACHE) -i -v
GO_TEST_FLAGS  ?= -v -cover
GO_PACKAGES    := $(shell go list ./... | grep -v vendor)

.PHONY: fmt
fmt:
	@go fmt $(GO_PACKAGES)

.PHONY: test
test: fmt
	@go test $(GO_TEST_FLAGS) $(GO_PACKAGES)

.PHONY: test-with-integration
test-with-integration: fmt
	@go test $(GO_TEST_FLAGS) -tags=integration $(GO_PACKAGES)

.PHONY: docker
docker:
	@docker build --build-arg db_version=$(DB_VERSION) --build-arg api_version=$(API_VERSION) --build-arg srv_version=$(SRV_VERSION) -f $(SERVER_DOCKERFILE) -t $(SERVER_IMAGE):$(IMAGE_VERSION) .
	@docker tag $(SERVER_IMAGE):$(IMAGE_VERSION) $(SERVER_IMAGE):latest
	@docker image prune -f --filter label=stage=server-intermediate
	@docker push $(SERVER_IMAGE)
.PHONY: push
push:
	@docker push $(SERVER_IMAGE)

.PHONY: protobuf
protobuf:
	@$(GENERATOR) \
	--go_out=plugins=grpc:. \
	--grpc-gateway_out=logtostderr=true:. \
	--gorm_out="engine=postgres:." \
	--swagger_out="atlas_patch=true:." \
	--atlas-query-validate_out=. \
	--atlas-validate_out="." \
	--validate_out="lang=go:." 	$(PROJECT_ROOT)/pkg/pb/cmdb.proto

.PHONY: vendor
vendor:
	@dep ensure -vendor-only

.PHONY: vendor-update
vendor-update:
	@dep ensure

.PHONY: clean
clean:
	@docker rmi -f $(shell docker images -q $(SERVER_IMAGE)) || true

.PHONY: up
up:
	cd repo/cmdb; helm install -f minikube.yaml -n cmdb .

.PHONY: down
down:
	helm delete cmdb


.PHONY: migrate-up
migrate-up:
	@migrate -database 'postgres://$(DATABASE_HOST)/cmdb?sslmode=disable' -path ./db/migrations up
#	@migrate -database 'postgres://$(DATABASE_HOST)/cmdb?sslmode=disable' -path ./db/migrations force 6

.PHONY: migrate-goto
migrate-goto:
	@migrate -database 'postgres://$(DATABASE_HOST)/cmdb?sslmode=disable' -path ./db/migrations goto 13

.PHONY: migrate-down
migrate-down:
	@migrate -database 'postgres://$(DATABASE_HOST)/cmdb?sslmode=disable' -path ./db/migrations down

.PHONY: erd
erd:
	@go-erd -path model |dot -Tsvg > doc/db/out.html
	@go-erd -path model |dot -Tpdf > doc/db/out.pdf
	@echo 'Create file://doc/db/out.html and file://doc/db/out.pdf'
