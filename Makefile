# Makefile for webmconv

# Variables
BINARY_NAME=webmconv
BUILD_DIR=build

# Default target
.PHONY: all
all: build

# Build the program
.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

# Run tests
.PHONY: test
test:
	go test ./...

# Run the program with arguments
.PHONY: run
run:
	go run main.go $(ARGS)

# Clean
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Install the program
.PHONY: install
install:
	go install .

# Install the CLI to the system
.PHONY: install-cli
install-cli: build
	cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@echo "CLI installed to /usr/local/bin/$(BINARY_NAME)"

# Help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build      - Build the program"
	@echo "  make test       - Run tests"
	@echo "  make run        - Run the program (use ARGS=\"...\" to pass arguments)"
	@echo "  make clean      - Remove generated files"
	@echo "  make install    - Install the program"
	@echo "  make install-cli - Build and copy executable to /usr/local/bin"
	@echo "  make help       - Show this message"