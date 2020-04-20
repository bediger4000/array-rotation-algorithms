package rotations

// Rotate uses the first k spaces of the array as
// a temp space, and swaps items from the first k
// spaces to their ultimate destination(s). Because
// of this, it can do fewer swaps than the cyclic
// rotation algorithm.
func Rotate(array []string, k int) (swaps int) {

	n := len(array)

	if k == 0 || n == k {
		return
	}

	// Swap each item to a spot k ahead of where it is,
	// but use (i+1)%k as the source. The first 0:k values
	// come from where they were originally, but k:k+k values
	// got swapped to 0:k indexes, and then get swapped out of
	// 0:k indexes
	finalDest := 0
	for i, j := 0, k; j < n; i, j = (i+1)%k, (j + 1) {
		array[i], array[j] = array[j], array[i]
		finalDest = i
		swaps++
	}

	moved := n - k
	untouched := n - 2*moved

	// All items got moved, exactly. No items in the temp area
	// need to be moved elsewhere, either.
	if untouched == 0 {
		return
	}

	// Some items left untouched. they have to be rotated around.
	if untouched > 0 {
		l := moved + untouched
		swaps += Rotate(array[:l], untouched)
		return
	}

	// No items left untouched, but some items in the temp area
	// need to be moved to their final destination. Because we
	// used the temp area in indexed order (0, 1, 2, ..., k-1)
	// all we need to do is rotate some of the temp area.
	rot := (n-1+k)%n - finalDest
	swaps += Rotate(array[:k], rot)
	return
}
