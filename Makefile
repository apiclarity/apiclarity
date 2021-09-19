e_Y=\033[1;33m
C_C=\033[0;36m
C_M=\033[0;35m
C_R=\033[0;41m
C_N=\033[0m
SHELL=/bin/bash

# Project variables
BINARY_NAME ?= apiclarity
DOCKER_REGISTRY ?= ghcr.io/apiclarity
VERSION ?= $(shell git rev-parse HEAD)
DOCKER_IMAGE ?= $(DOCKER_REGISTRY)/$(BINARY_NAME)
DOCKER_TAG ?= ${VERSION}

# Dependency versions
GOLANGCI_VERSION = 1.42.0
LICENSEI_VERSION = 0.3.1

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

.PHONY: ui
ui: ## Build UI
	@(echo "Building UI ..." )
	@(cd ui; npm i ; npm run build; )
	@ls -l ui/build  

.PHONY: backend
backend: ## Build Backend
	@(echo "Building Backend ..." )
	@(cd backend && go build -o bin/backend cmd/backend/main.go && ls -l bin/)

.PHONY: backend_linux
backend_linux: ## Build Backend Linux
	@(echo "Building Backend linux..." )
	@(cd backend && GOOS=linux go build -o bin/backend_linux cmd/backend/main.go && ls -l bin/)

.PHONY: backend_test
backend_test: ## Build Backend test
	@(echo "Building Backend test ..." )
	@(cd backend && go build -o bin/backend_test cmd/test/main.go && ls -l bin/)

.PHONY: api
api: ## Generating API code
	@(echo "Generating API code ..." )
	@(cd api; ./generate.sh)

.PHONY: docker
docker: ## Build Docker image 
	@(echo "Building docker image ..." )
	docker build --build-arg VERSION=${VERSION} \
		--build-arg BUILD_TIMESTAMP=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ") \
		--build-arg COMMIT_HASH=$(shell git rev-parse HEAD) \
		-t ${DOCKER_IMAGE}:${DOCKER_TAG} .

.PHONY: push-docker
push-docker: docker ## Build and Push Docker image
	@echo "Publishing Docker image ..."
	docker push ${DOCKER_IMAGE}:${DOCKER_TAG}

.PHONY: test
test: ## Run Unit Tests
	@(cd backend && FAKE_DATA=true go test ./pkg/...)

.PHONY: clean
clean: clean-ui clean-backend ## Clean all build artifacts

.PHONY: clean-ui
clean-ui: 
	@(rm -rf ui/build ; echo "UI cleanup done" )

.PHONY: clean-backend
clean-backend: 
	@(rm -rf bin ; echo "Backend cleanup done" )

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint
bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ./bin/ v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

.PHONY: lint
lint: bin/golangci-lint ## Run linter
	cd backend && ../bin/golangci-lint run

.PHONY: fix
fix: bin/golangci-lint ## Fix lint violations
	cd backend && ../bin/golangci-lint run --fix

bin/licensei: bin/licensei-${LICENSEI_VERSION}
	@ln -sf licensei-${LICENSEI_VERSION} bin/licensei
bin/licensei-${LICENSEI_VERSION}:
	@mkdir -p bin
	curl -sfL https://raw.githubusercontent.com/goph/licensei/master/install.sh | bash -s v${LICENSEI_VERSION}
	@mv bin/licensei $@

.PHONY: license-check
license-check: bin/licensei ## Run license check
	bin/licensei header
	cd backend && ../bin/licensei check --config=../.licensei.toml

.PHONY: license-cache
license-cache: bin/licensei ## Generate license cache
	bin/licensei cache

.PHONY: check
check: lint test ## Run tests and linters