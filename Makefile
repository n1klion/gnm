# Makefile for building the application

# Set the binary name
BINARY_NAME=gnm

# Set the Go build command
GOBUILD=go build

# Build directory
BUILD_DIR=./build

# Source file
SOURCE_FILE=./main.go

.PHONY: all build clean

all: build

build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_FILE)

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)