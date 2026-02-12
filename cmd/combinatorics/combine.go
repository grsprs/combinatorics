package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/grsprs/combinatorics/combinations"
)

// handleCombine processes the combine command.
// Generates all k-combinations of the provided items.
// Requires -k flag to specify combination size.
// Supports --json flag for JSON output format.
func handleCombine(args []string) {
	if len(args) < 3 {
		fmt.Fprintln(os.Stderr, "Error: combine requires -k flag and items")
		fmt.Fprintln(os.Stderr, "Usage: combinatorics combine -k <k> <items...>")
		os.Exit(1)
	}

	if args[0] != "-k" {
		fmt.Fprintln(os.Stderr, "Error: -k flag is required")
		fmt.Fprintln(os.Stderr, "Usage: combinatorics combine -k <k> <items...>")
		os.Exit(1)
	}

	k, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid k value '%s'\n", args[1])
		os.Exit(1)
	}

	items := filterFlags(args[2:])
	if len(items) == 0 {
		fmt.Fprintln(os.Stderr, "Error: no items provided")
		os.Exit(1)
	}

	combs, err := combinations.Combine(items, k)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if hasFlag(args, "--json") {
		output := map[string]interface{}{
			"command":      "combine",
			"k":            k,
			"items":        items,
			"combinations": combs,
			"count":        len(combs),
		}
		jsonData, _ := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(jsonData))
	} else {
		fmt.Printf("%d-combinations of %v:\n", k, items)
		for i, comb := range combs {
			fmt.Printf("%d: %v\n", i+1, comb)
		}
		fmt.Printf("\nTotal: %d combinations\n", len(combs))
	}
}
