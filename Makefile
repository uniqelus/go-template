PROJECT_NAME = $(shell basename $(shell pwd))
SERVICES := $(shell ls cmd)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0

# Version information
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