package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grsprs/combinatorics/permutations"
)

// handlePermute processes the permute command.
// Generates all permutations of the provided items.
// Supports --json flag for JSON output format.
func handlePermute(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Error: permute requires at least 1 item")
		fmt.Fprintln(os.Stderr, "Usage: combinatorics permute <items...>")
		os.Exit(1)
	}

	items := filterFlags(args)
	if len(items) == 0 {
		fmt.Fprintln(os.Stderr, "Error: no items provided")
		os.Exit(1)
	}

	perms, err := permutations.Permute(items)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if hasFlag(args, "--json") {
		output := map[string]interface{}{
			"command":      "permute",
			"items":        items,
			"permutations": perms,
			"count":        len(perms),
		}
		jsonData, _ := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(jsonData))
	} else {
		fmt.Printf("Permutations of %v:\n", items)
		for i, perm := range perms {
			fmt.Printf("%d: %v\n", i+1, perm)
		}
		fmt.Printf("\nTotal: %d permutations\n", len(perms))
	}
}
