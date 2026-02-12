# Combinatorics

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![CI Status](https://github.com/grsprs/combinatorics/workflows/CI/badge.svg)](https://github.com/grsprs/combinatorics/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/grsprs/combinatorics)](https://goreportcard.com/report/github.com/grsprs/combinatorics)
[![codecov](https://codecov.io/gh/grsprs/combinatorics/branch/main/graph/badge.svg)](https://codecov.io/gh/grsprs/combinatorics)
[![Go Reference](https://pkg.go.dev/badge/github.com/grsprs/combinatorics.svg)](https://pkg.go.dev/github.com/grsprs/combinatorics)

Production-grade combinatorics library in Go featuring efficient implementations of permutations, combinations, factorials, and binomial coefficients with generics support and CLI tool.

## Features

- **Factorial** - Efficient computation using `math/big` for large values
- **Binomial Coefficient** - Optimized nCr without full factorial expansion
- **Permutations** - Generate all permutations with Heap's algorithm
- **Combinations** - Generate k-combinations efficiently
- **Iterators** - Memory-efficient lazy generation using channels
- **Generics** - Type-safe operations for any comparable type
- **CLI Tool** - Command-line interface for quick calculations
- **100% Test Coverage** - Comprehensive test suite with benchmarks

## Installation

### As a Library

```bash
go get github.com/grsprs/combinatorics
```

### As a CLI Tool

```bash
go install github.com/grsprs/combinatorics/cmd/combinatorics@latest
```

Or download pre-built binaries from [Releases](https://github.com/grsprs/combinatorics/releases).

## Quick Start

### Library Usage

### Factorial

```go
import "github.com/grsprs/combinatorics/factorial"

result, err := factorial.Factorial(10)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // 3628800
```

### Binomial Coefficient

```go
import "github.com/grsprs/combinatorics/binomial"

// Calculate C(10, 3) = 120
result, err := binomial.Binomial(10, 3)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // 120
```

### Permutations

```go
import "github.com/grsprs/combinatorics/permutations"

// Generate all permutations
items := []int{1, 2, 3}
perms, err := permutations.Permute(items)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Total: %d\n", len(perms)) // 6

// Lazy generation with iterator
for perm := range permutations.PermuteIter(items) {
    fmt.Println(perm)
}
```

### Combinations

```go
import "github.com/grsprs/combinatorics/combinations"

// Generate all 2-combinations
items := []int{1, 2, 3, 4}
combs, err := combinations.Combine(items, 2)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Total: %d\n", len(combs)) // 6

// Lazy generation with iterator
for comb := range combinations.CombineIter(items, 2) {
    fmt.Println(comb)
}
```

### CLI Usage

```bash
# Calculate factorial
combinatorics factorial 10
# Output: 10! = 3628800

# Calculate binomial coefficient
combinatorics binomial 10 3
# Output: C(10, 3) = 120

# Generate permutations
combinatorics permute 1 2 3
# Output: All 6 permutations

# Generate combinations
combinatorics combine -k 2 1 2 3 4
# Output: All 6 2-combinations

# JSON output
combinatorics factorial 10 --json
```

## API Documentation

### Factorial Package

```go
func Factorial(n int) (*big.Int, error)
```

Computes n! for non-negative integers. Returns error for negative input.

**Time Complexity:** O(n)  
**Space Complexity:** O(1)

### Binomial Package

```go
func Binomial(n, k int) (*big.Int, error)
```

Computes C(n, k) = n! / (k! × (n-k)!) without full factorial expansion.

**Time Complexity:** O(k)  
**Space Complexity:** O(1)

### Permutations Package

```go
func Permute[T any](items []T) ([][]T, error)
func PermuteIter[T any](items []T) <-chan []T
```

Generates all permutations using Heap's algorithm.

**Time Complexity:** O(n!)  
**Space Complexity:** O(n! × n) for `Permute`, O(n) for `PermuteIter`

### Combinations Package

```go
func Combine[T any](items []T, k int) ([][]T, error)
func CombineIter[T any](items []T, k int) <-chan []T
```

Generates all k-combinations.

**Time Complexity:** O(C(n, k))  
**Space Complexity:** O(C(n, k) × k) for `Combine`, O(k) for `CombineIter`

## Performance

Benchmarks run on Intel Core i7-4790 @ 3.60GHz:

### Factorial
```
BenchmarkFactorial/n=10      4,424,268 ops    274.2 ns/op    88 B/op    3 allocs/op
BenchmarkFactorial/n=20      2,328,636 ops    448.2 ns/op    88 B/op    3 allocs/op
BenchmarkFactorial/n=50      1,000,000 ops   1023   ns/op    88 B/op    3 allocs/op
BenchmarkFactorial/n=100       595,774 ops   2189   ns/op   184 B/op    4 allocs/op
```

### Binomial
```
BenchmarkBinomial/C(10,5)    3,626,236 ops    330.7 ns/op    88 B/op    3 allocs/op
BenchmarkBinomial/C(20,10)   2,181,832 ops    539.3 ns/op    88 B/op    3 allocs/op
BenchmarkBinomial/C(50,25)   1,000,000 ops   1247   ns/op    88 B/op    3 allocs/op
BenchmarkBinomial/C(100,50)    336,013 ops   3768   ns/op    88 B/op    3 allocs/op
```

### Permutations
```
BenchmarkPermute/n=5           104,766 ops    11,109 ns/op    13,288 B/op    128 allocs/op
BenchmarkPermute/n=7             2,499 ops   524,873 ns/op   832,491 B/op  5,056 allocs/op
BenchmarkPermuteIter/n=5        22,374 ops    54,267 ns/op     5,921 B/op    122 allocs/op
BenchmarkPermuteIter/n=7           508 ops 2,272,658 ns/op   322,722 B/op  5,042 allocs/op
```

### Combinations
```
BenchmarkCombine/C(10,5)        52,717 ops    20,393 ns/op    27,864 B/op    262 allocs/op
BenchmarkCombine/C(15,7)         2,205 ops   584,083 ns/op   921,835 B/op  6,452 allocs/op
BenchmarkCombineIter/C(10,5)    10,000 ops   111,432 ns/op    12,322 B/op    255 allocs/op
BenchmarkCombineIter/C(15,7)       415 ops 2,945,400 ns/op   412,083 B/op  6,438 allocs/op
```

## Error Handling

All functions return descriptive errors:

```go
// Factorial errors
factorial.ErrNegativeInput  // n < 0

// Binomial errors
binomial.ErrNegativeN       // n < 0
binomial.ErrNegativeK       // k < 0
binomial.ErrKGreaterN       // k > n

// Permutations errors
permutations.ErrEmptyInput  // nil input

// Combinations errors
combinations.ErrEmptyInput  // nil input
combinations.ErrNegativeK   // k < 0
combinations.ErrKGreaterN   // k > n
```

## Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run with race detector
go test -race ./...

# Run benchmarks
go test -bench=. -benchmem ./...
```

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) for details.

## Author

**Spiros Nikoloudakis** ([@grsprs](https://github.com/grsprs))  
Email: sp.nikoloudakis@gmail.com

## Acknowledgments

- Heap's algorithm for efficient permutation generation
- Optimized binomial coefficient computation
- Go generics for type-safe operations

## Project Status

✅ Production-ready  
✅ 100% test coverage  
✅ Comprehensive benchmarks  
✅ Full documentation

---

**Copyright © 2026 Spiros Nikoloudakis**
