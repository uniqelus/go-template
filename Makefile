PROJECT_NAME = $(shell basename $(shell pwd))
SERVICES := $(shell ls cmd)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS = -X go-template/pkg/version.Version=$(VERSION) \
          -X go-template/pkg/version.GitCommit=$(GIT_COMMIT) \
          -X go-template/pkg/version.BuildTime=$(BUILD_TIME)

init:
	rm -f go.mod go.sum
	go mod init $(PROJECT_NAME)
	find . -type f -name "*.go" -print0 | xargs -0 sed -i "s|go-template|$(MODULE)|g"
	go mod tidy

binaries: $(addsuffix -binary, $(SERVICES))
%-binary:
	@mkdir -p bin
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" $(FLAGS) -o bin/$* ./cmd/$*

lint:
	golangci-lint run $(if $(FIX),,--fix) ./...

.PHONY: version
version:
	@echo "Version: $(VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"

EXCLUDED_FROM_DOCKER ?= ""
DOCKER_SERVICES := $(filter-out $(EXCLUDED_FROM_DOCKER), $(SERVICES))

DOCKERFILES_DIR := deployments/docker

images: $(addsuffix -image, $(DOCKER_SERVICES))
%-image: %-binary
	DOCKER_BUILDKIT=1 docker build \
		-t $*:latest \
		-f $(DOCKERFILES_DIR)/service.dockerfile \
		--build-arg SERVICE_NAME=$* .

ENV_FILE = .env
DOCKER_COMPOSE_FILE = deployments/docker/docker-compose.yaml
DOCKER_COMPOSE = docker compose -p $(PROJECT_NAME) -f $(DOCKER_COMPOSE_FILE) --env-file $(ENV_FILE)

up: images
	$(DOCKER_COMPOSE) up --detach

down:
	$(DOCKER_COMPOSE) down --remove-orphans -v

ps:
	$(DOCKER_COMPOSE) ps --all --format "table {{.Service}}\t{{.Status}}\t{{.Ports}}"

start: $(addsuffix -start, $(DOCKER_SERVICES))
%-start:
	$(DOCKER_COMPOSE) start $*

stop: $(addsuffix -stop, $(DOCKER_SERVICES))
%-stop:
	$(DOCKER_COMPOSE) stop $*

restart: $(addsuffix -restart, $(DOCKER_SERVICES))
%-restart:
	$(DOCKER_COMPOSE) restart $*