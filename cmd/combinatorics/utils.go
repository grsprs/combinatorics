package main

// hasFlag checks if a specific flag exists in the arguments.
func hasFlag(args []string, flag string) bool {
	for _, arg := range args {
		if arg == flag {
			return true
		}
	}
	return false
}

// filterFlags removes flag arguments from the list, returning only data arguments.
func filterFlags(args []string) []string {
	var result []string
	for _, arg := range args {
		if arg != "--json" {
			result = append(result, arg)
		}
	}
	return result
}
