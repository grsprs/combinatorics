# Contributing to Combinatorics

## Prerequisites

- Go 1.22 or higher
- Git

## Development Setup

```bash
# Clone the repository
git clone https://github.com/grsprs/combinatorics.git
cd combinatorics

# Verify Go installation
go version

# Download dependencies
go mod download
```

## Building

```bash
# Build all packages
go build ./...
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run tests with race detector
go test -race ./...

# Run benchmarks
go test -bench=. ./...
```

## Code Quality

```bash
# Format code
go fmt ./...

# Run linter (requires golangci-lint)
golangci-lint run

# Run security scanner (requires gosec)
gosec ./...

# Check for vulnerabilities (requires govulncheck)
govulncheck ./...
```

## Contribution Guidelines

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Write tests for your changes
4. Ensure all tests pass and coverage meets requirements (80%+)
5. Format your code (`go fmt ./...`)
6. Commit with conventional commit messages
7. Push to your fork and submit a pull request

## Commit Message Format

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>: <description>

[optional body]

[optional footer]
```

Types: `feat`, `fix`, `docs`, `test`, `refactor`, `perf`, `chore`

## Pull Request Process

1. Update documentation if needed
2. Add tests for new functionality
3. Ensure CI passes
4. Request review from maintainers
5. Address review feedback

## Code Standards

- Follow idiomatic Go conventions
- Write clear, self-documenting code
- Add GoDoc comments for all exported identifiers
- Maintain test coverage above 80%
- Include benchmarks for performance-critical code
