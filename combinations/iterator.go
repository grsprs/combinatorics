package combinations

// CombineIter generates k-combinations lazily using a channel.
// Useful for large C(n, k) to avoid materializing all combinations in memory.
// The channel is closed when all combinations are generated.
func CombineIter[T any](items []T, k int) <-chan []T {
	ch := make(chan []T)

	go func() {
		defer close(ch)

		if items == nil {
			return
		}

		n := len(items)

		if k < 0 || k > n {
			return
		}

		if k == 0 {
			ch <- []T{}
			return
		}

		if k == n {
			comb := make([]T, n)
			copy(comb, items)
			ch <- comb
			return
		}

		indices := make([]int, k)
		for i := range indices {
			indices[i] = i
		}

		for {
			comb := make([]T, k)
			for i, idx := range indices {
				comb[i] = items[idx]
			}
			ch <- comb

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
	}()

	return ch
}
