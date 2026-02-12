package combinations

import "fmt"

var (
	ErrEmptyInput = fmt.Errorf("combinations: input slice cannot be nil")
	ErrNegativeK  = fmt.Errorf("combinations: k must be non-negative")
	ErrKGreaterN  = fmt.Errorf("combinations: k cannot be greater than n")
)

// Combine generates all k-combinations of the input slice.
// Returns a slice of all combinations.
// Time complexity: O(C(n, k))
// Space complexity: O(C(n, k) * k)
func Combine[T any](items []T, k int) ([][]T, error) {
	if items == nil {
		return nil, ErrEmptyInput
	}

	n := len(items)

	if k < 0 {
		return nil, fmt.Errorf("%w: got %d", ErrNegativeK, k)
	}

	if k > n {
		return nil, fmt.Errorf("%w: k=%d, n=%d", ErrKGreaterN, k, n)
	}

	if k == 0 {
		return [][]T{{}}, nil
	}

	if k == n {
		return [][]T{append([]T{}, items...)}, nil
	}

	var result [][]T
	indices := make([]int, k)
	for i := range indices {
		indices[i] = i
	}

	for {
		comb := make([]T, k)
		for i, idx := range indices {
			comb[i] = items[idx]
		}
		result = append(result, comb)

		// Find the rightmost index that can be incremented
		i := k - 1
		for i >= 0 && indices[i] == n-k+i {
			i--
		}

		if i < 0 {
			break
		}

		indices[i]++
		for j := i + 1; j < k; j++ {
			indices[j] = indices[j-1] + 1
		}
	}

	return result, nil
}
