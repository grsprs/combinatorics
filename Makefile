.PHONY: test coverage lint security build clean install help

help:
	@echo "Available targets:"
	@echo "  test      - Run all tests"
	@echo "  coverage  - Run tests with coverage report"
	@echo "  lint      - Run linter"
	@echo "  security  - Run security scans"
	@echo "  build     - Build CLI binary"
	@echo "  install   - Install CLI tool"
	@echo "  clean     - Remove build artifacts"

test:
	go test -v -race ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

lint:
	golangci-lint run

security:
	gosec ./...
	govulncheck ./...

build:
	go build -o bin/combinatorics.exe ./cmd/combinatorics

install:
	go install ./cmd/combinatorics

clean:
	rm -rf bin/ coverage.out coverage.txt *.exe

all: test lint security build
