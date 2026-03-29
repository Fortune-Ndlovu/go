# Set the default target that runs when `make` is executed with no args
.DEFAULT_GOAL := build

# Format all Go source files in the project
fmt:
	go fmt ./...
.PHONY: fmt  # Marks this target as not a real file

# Run linter after formatting (depends on fmt)
lint: fmt
	golint ./...
.PHONY: lint

# Run static analysis checks after formatting (depends on fmt)
vet: fmt
	go vet ./...
.PHONY: vet

# Build the project after running vet (which already formats)
build: vet
	go build ./...
.PHONY: build
