# Makefile for LogAnalyzer

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary name
BINARY_NAME=loganalyzer

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./main.go

# Run the application
run:
	$(GOCMD) run main.go

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Test the application
test:
	$(GOTEST) -v ./...

# Get dependencies
deps:
	$(GOGET) -v ./...

# Default target
all: clean build

.PHONY: build run clean test deps all