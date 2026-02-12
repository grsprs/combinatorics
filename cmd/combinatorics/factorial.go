package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/grsprs/combinatorics/factorial"
)

// handleFactorial processes the factorial command.
// Parses the input argument, validates it, and computes n!.
// Supports --json flag for JSON output format.
func handleFactorial(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Error: factorial requires 1 argument")
		fmt.Fprintln(os.Stderr, "Usage: combinatorics factorial <n>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		// #nosec G705 - CLI output, not web context
		fmt.Fprintf(os.Stderr, "Error: invalid number '%s'\n", args[0])
		os.Exit(1)
	}

	result, err := factorial.Factorial(n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if hasFlag(args, "--json") {
		output := map[string]interface{}{
			"command": "factorial",
			"input":   n,
			"result":  result.String(),
		}
		jsonData, _ := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(jsonData))
	} else {
		fmt.Printf("%d! = %s\n", n, result)
	}
}
