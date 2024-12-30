# Variables
APP_NAME := githubAgent
SRC_FILE := cmd/server/main.go
OUTPUT_DIR := bin
OUTPUT_FILE := $(OUTPUT_DIR)/$(APP_NAME)
GO := go

# Default target
.PHONY: all
all: lint build

# Lint the Go source code
.PHONY: lint
lint:
	@echo "Running linter on $(SRC_FILE)..."
	#golangci-lint run ./... || echo "Install golangci-lint: https://golangci-lint.run/usage/install/"

# Build the Go application
.PHONY: build
build:
	@echo "Building $(SRC_FILE)..."
	@mkdir -p $(OUTPUT_DIR)
	$(GO) build -o $(OUTPUT_FILE) $(SRC_FILE)
	@echo "Build complete: $(OUTPUT_FILE)"

# Run the Go application
.PHONY: run
run:
	@echo "Running $(OUTPUT_FILE)..."
	$(OUTPUT_FILE)

# Run Go tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test ./...

# Clean the build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(OUTPUT_DIR)
	@echo "Clean complete."

# Format the code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Tidy the Go module dependencies
.PHONY: tidy
tidy:
	@echo "Tidying Go modules..."
	$(GO) mod tidy

# Help message
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make lint       Run linter on the source code"
	@echo "  make build      Build the application"
	@echo "  make run        Run the application"
	@echo "  make test       Run all tests"
	@echo "  make clean      Remove build artifacts"
	@echo "  make fmt        Format the Go source code"
	@echo "  make tidy       Tidy up Go modules"
	@echo "  make help       Show this help message"
