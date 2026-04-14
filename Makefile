PROJECT_NAME = $(shell basename $(shell pwd))
SERVICES := $(shell ls cmd)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0

init:
	rm -f go.mod go.sum
	go mod init $(PROJECT_NAME)
	find . -type f -name "*.go" -print0 | xargs -0 sed -i "s|go-template|$(MODULE)|g"
	go mod tidy

binaries: $(addsuffix -binary, $(SERVICES))
%-binary:
	@mkdir -p bin
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(FLAGS) -o bin/$* ./cmd/$*