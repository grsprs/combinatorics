package factorial

import (
	"fmt"
	"math/big"
)

var ErrNegativeInput = fmt.Errorf("factorial: input must be non-negative")

// Factorial computes n! for non-negative integers.
// Returns error for negative input.
// Uses big.Int to handle large values without overflow.
func Factorial(n int) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("%w: got %d", ErrNegativeInput, n)
	}

	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result, nil
}
