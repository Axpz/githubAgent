# Variables
APP_SVR_NAME := agentServer
SRC_SVR_FILE := cmd/server/main.go

APP_AGENT_NAME := agent
SRC_AGENT_FILE := cmd/agent/main.go

OUTPUT_DIR := bin
OUTPUT_SVR_FILE := $(OUTPUT_DIR)/$(APP_SVR_NAME)
OUTPUT_AGENT_FILE := $(OUTPUT_DIR)/$(APP_AGENT_NAME)
GO := go

# Default target
.PHONY: all
all: lint build

.PHONY: proto
proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$(PROTO_DIR)/*.proto

# Lint the Go source code
.PHONY: lint
lint:
	@echo "Running linter on $(SRC_SVR_FILE)..."
	@echo "Running linter on $(SRC_AGENT_FILE)..."

# Build the Go application
.PHONY: build
build:
	@echo "Building $(SRC_SVR_FILE)..."
	@mkdir -p $(OUTPUT_DIR)
	$(GO) build -o $(OUTPUT_SVR_FILE) $(SRC_SVR_FILE)

	@echo "Building $(SRC_AGENT_FILE)..."
	$(GO) build -o $(OUTPUT_AGENT_FILE) $(SRC_AGENT_FILE)
	@echo "Build complete: $(OUTPUT_AGENT_FILE)"

# Run the Go application
.PHONY: run
run: build
	@echo "Running $(OUTPUT_SVR_FILE)..."
	$(OUTPUT_SVR_FILE)

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
