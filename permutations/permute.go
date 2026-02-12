package permutations

import "fmt"

var ErrEmptyInput = fmt.Errorf("permutations: input slice cannot be nil")

// Permute generates all permutations of the input slice.
// Returns a slice of all permutations.
// Time complexity: O(n!)
// Space complexity: O(n! * n)
func Permute[T any](items []T) ([][]T, error) {
	if items == nil {
		return nil, ErrEmptyInput
	}

	if len(items) == 0 {
		return [][]T{{}}, nil
	}

	var result [][]T
	permute(items, 0, &result)
	return result, nil
}

// permute uses Heap's algorithm for in-place permutation generation
func permute[T any](items []T, start int, result *[][]T) {
	if start == len(items)-1 {
		perm := make([]T, len(items))
		copy(perm, items)
		*result = append(*result, perm)
		return
	}

	for i := start; i < len(items); i++ {
		items[start], items[i] = items[i], items[start]
		permute(items, start+1, result)
		items[start], items[i] = items[i], items[start]
	}
}
