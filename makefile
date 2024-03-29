GO := go
BINARY_NAME := dondu
BUILD_DIR := build
DB_DIR := database

build:
	@cd UI && bun vite build
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) .

build-server:
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) .

build-web:
	@cd UI && bun vite build

dev:
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) .
	$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@rm -rf $(BUILD_DIR)
clean-all:
	@rm -rf $(BUILD_DIR)
	@rm -rf $(DB_DIR)

.PHONY: build build-server build-web dev clean clean-all
