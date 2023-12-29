PACKAGES := $(shell go list ./...)
name := $(shell basename ${PWD})

# avoid that in normal mode, make will asume that the target name is the name of file.
.PHONY: init all build run dev test clean

# Build the application
all: build

init:
	go mod init ${module}
	go install github.com/cosmtrek/air@latest

build:
	@echo "Building..."
	@templ generate
	@go build -o bin/main main.go

# Run the application
run:
	@templ generate
	@go run .

# dev the application
dev:
	air

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f bin/main


