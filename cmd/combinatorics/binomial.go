package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/grsprs/combinatorics/binomial"
)

// handleBinomial processes the binomial coefficient command.
// Parses n and k arguments, validates them, and computes C(n, k).
// Supports --json flag for JSON output format.
func handleBinomial(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: binomial requires 2 arguments")
		fmt.Fprintln(os.Stderr, "Usage: combinatorics binomial <n> <k>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		// #nosec G705 - CLI output, not web context
		fmt.Fprintf(os.Stderr, "Error: invalid number '%s'\n", args[0])
		os.Exit(1)
	}

	k, err := strconv.Atoi(args[1])
	if err != nil {
		// #nosec G705 - CLI output, not web context
		fmt.Fprintf(os.Stderr, "Error: invalid number '%s'\n", args[1])
		os.Exit(1)
	}

	result, err := binomial.Binomial(n, k)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if hasFlag(args, "--json") {
		output := map[string]interface{}{
			"command": "binomial",
			"n":       n,
			"k":       k,
			"result":  result.String(),
		}
		jsonData, _ := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(jsonData))
	} else {
		fmt.Printf("C(%d, %d) = %s\n", n, k, result)
	}
}
