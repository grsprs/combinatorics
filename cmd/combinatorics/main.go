package main

import (
	"fmt"
	"os"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "factorial", "fact":
		handleFactorial(os.Args[2:])
	case "binomial", "binom":
		handleBinomial(os.Args[2:])
	case "permute", "perm":
		handlePermute(os.Args[2:])
	case "combine", "comb":
		handleCombine(os.Args[2:])
	case "version", "-v", "--version":
		fmt.Printf("combinatorics v%s\n", version)
	case "help", "-h", "--help":
		printUsage()
	default:
		// #nosec G705 - CLI output, not web context
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

// printUsage displays the CLI help message with available commands and examples.
func printUsage() {
	fmt.Println("Combinatorics CLI - Production-grade combinatorics calculations")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  combinatorics <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  factorial <n>              Calculate n!")
	fmt.Println("  binomial <n> <k>           Calculate C(n, k)")
	fmt.Println("  permute <items...>         Generate all permutations")
	fmt.Println("  combine -k <k> <items...>  Generate k-combinations")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help                 Show this help message")
	fmt.Println("  -v, --version              Show version")
	fmt.Println("  --json                     Output in JSON format")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  combinatorics factorial 10")
	fmt.Println("  combinatorics binomial 10 3")
	fmt.Println("  combinatorics permute 1 2 3")
	fmt.Println("  combinatorics combine -k 2 1 2 3 4")
}
