# Go parameters
BINARY_NAME=myapp
MAIN_FILE=main.go

.PHONY: all build run test clean lint fmt tidy help

all: build

## Build the project
build:
	@echo "🔨 Building..."
	go build -o bin/$(BINARY_NAME) $(MAIN_FILE)

## Run the project
run:
	@echo "🚀 Running..."
	go run $(MAIN_FILE)

## Run tests with coverage
test:
	@echo "🧪 Running tests..."
	go test ./... -v -coverprofile=coverage.out

## Clean build cache & binaries
clean:
	@echo "🧹 Cleaning..."
	go clean
	rm -rf bin coverage.out

## Lint the code using golangci-lint (install if not present)
lint:
	@echo "🔍 Linting..."
	golangci-lint run ./...

## Format the code
fmt:
	@echo "✨ Formatting..."
	go fmt ./...

## Ensure modules are tidy
tidy:
	@echo "📦 Tidying modules..."
	go mod tidy

## Show help
help:
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'
