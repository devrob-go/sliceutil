# Makefile for SliceUtil Go package
# Provides targets for building, testing, linting, and development

.PHONY: help build test test-coverage test-benchmark clean lint format check-fmt vet run-example install-deps update-deps

# Default target
help:
	@echo "SliceUtil - Go Package for Slice Operations"
	@echo ""
	@echo "Available targets:"
	@echo "  build          - Build the package"
	@echo "  test           - Run all tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  test-benchmark - Run benchmark tests"
	@echo "  lint           - Run golangci-lint"
	@echo "  format         - Format Go code"
	@echo "  check-fmt      - Check if code is formatted"
	@echo "  vet            - Run go vet"
	@echo "  clean          - Clean build artifacts"
	@echo "  run-example    - Run the example program"
	@echo "  install-deps   - Install development dependencies"
	@echo "  update-deps    - Update Go dependencies"
	@echo "  help           - Show this help message"

# Build the package
build:
	@echo "Building SliceUtil package..."
	go build ./...

# Run all tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run benchmark tests
test-benchmark:
	@echo "Running benchmark tests..."
	go test -bench=. -benchmem ./...

# Run linting
lint:
	@echo "Running golangci-lint..."
	golangci-lint run

# Format Go code
format:
	@echo "Formatting Go code..."
	go fmt ./...
	gofmt -s -w .

# Check if code is formatted
check-fmt:
	@echo "Checking code formatting..."
	@if [ -n "$$(gofmt -l .)" ]; then \
		echo "Code is not formatted. Run 'make format' to fix."; \
		exit 1; \
	fi
	@echo "Code is properly formatted."

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	go clean
	rm -f coverage.out coverage.html
	rm -rf dist/

# Run the example program
run-example:
	@echo "Running example program..."
	go run cmd/example/main.go

# Install development dependencies
install-deps:
	@echo "Installing development dependencies..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest

# Update Go dependencies
update-deps:
	@echo "Updating Go dependencies..."
	go get -u ./...
	go mod tidy

# Run all checks (format, vet, test)
check: check-fmt vet test

# Build and test
all: build test

# Development workflow
dev: format vet test run-example

# CI/CD pipeline
ci: check-fmt vet test-coverage test-benchmark

# Release preparation
release: clean check-fmt vet test-coverage build
	@echo "Release preparation completed successfully!"

# Show package information
info:
	@echo "Package: github.com/devrob-go/sliceutil"
	@echo "Go version: $(shell go version)"
	@echo "Go modules: $(shell go env GOMOD)"
	@echo "Go workspace: $(shell go env GOWORK)"
	@echo ""
	@echo "Dependencies:"
	@go list -m all

# Show test coverage summary
coverage-summary:
	@echo "Test coverage summary:"
	@go test -cover ./... | grep -E "(coverage|PASS|FAIL)"

# Run specific test file
test-file:
	@if [ -z "$(FILE)" ]; then \
		echo "Usage: make test-file FILE=path/to/test.go"; \
		exit 1; \
	fi
	@echo "Running tests in $(FILE)..."
	go test -v $(FILE)

# Run tests with race detection
test-race:
	@echo "Running tests with race detection..."
	go test -race ./...

# Generate documentation
docs:
	@echo "Generating documentation..."
	godoc -http=:6060 &
	@echo "Documentation server started at http://localhost:6060"
	@echo "Press Ctrl+C to stop"

# Install package locally
install:
	@echo "Installing package locally..."
	go install ./...

# Uninstall package
uninstall:
	@echo "Uninstalling package..."
	go clean -i ./...

# Show help for specific target
help-%:
	@echo "Help for target '$*':"
	@make -n $* 2>/dev/null || echo "No help available for target '$*'"

# Default target when no arguments provided
.DEFAULT_GOAL := help
