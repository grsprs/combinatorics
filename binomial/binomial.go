package binomial

import (
	"fmt"
	"math/big"
)

var (
	ErrNegativeN = fmt.Errorf("binomial: n must be non-negative")
	ErrNegativeK = fmt.Errorf("binomial: k must be non-negative")
	ErrKGreaterN = fmt.Errorf("binomial: k cannot be greater than n")
)

// Binomial computes the binomial coefficient C(n, k) = n! / (k! * (n-k)!)
// Uses optimization to avoid full factorial expansion.
// Returns error if n < 0, k < 0, or k > n.
func Binomial(n, k int) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("%w: got %d", ErrNegativeN, n)
	}
	if k < 0 {
		return nil, fmt.Errorf("%w: got %d", ErrNegativeK, k)
	}
	if k > n {
		return nil, fmt.Errorf("%w: k=%d, n=%d", ErrKGreaterN, k, n)
	}

	// Optimization: C(n, k) = C(n, n-k), use smaller k
	if k > n-k {
		k = n - k
	}

	if k == 0 {
		return big.NewInt(1), nil
	}

	result := big.NewInt(1)
	for i := 0; i < k; i++ {
		result.Mul(result, big.NewInt(int64(n-i)))
		result.Div(result, big.NewInt(int64(i+1)))
	}

	return result, nil
}
