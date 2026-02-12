package permutations

// PermuteIter generates permutations lazily using a channel.
// Useful for large n to avoid materializing all permutations in memory.
// The channel is closed when all permutations are generated.
func PermuteIter[T any](items []T) <-chan []T {
	ch := make(chan []T)

	go func() {
		defer close(ch)

		if items == nil || len(items) == 0 {
			if items != nil && len(items) == 0 {
				ch <- []T{}
			}
			return
		}

		permuteIter(items, 0, ch)
	}()

	return ch
}

func permuteIter[T any](items []T, start int, ch chan<- []T) {
	if start == len(items)-1 {
		perm := make([]T, len(items))
		copy(perm, items)
		ch <- perm
		return
	}

	for i := start; i < len(items); i++ {
		items[start], items[i] = items[i], items[start]
		permuteIter(items, start+1, ch)
		items[start], items[i] = items[i], items[start]
	}
}
