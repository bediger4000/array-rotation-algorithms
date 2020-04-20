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

// Slowrotate does k rotations by 1
func Slowrotate(array []string, k int) (swaps int) {
	// Rotate an array by 1 rotate number of times
	for i := 0; i < k; i++ {
		// Rotate an array by 1
		tmp := array[i]
		for j := i + 1; j != i; j = (j + 1) % len(array) {
			array[j], tmp = tmp, array[j]
			swaps++
		}
		array[i] = tmp
	}
	return
}

// Cyclic rotates an array k items to the right by jumping each item k
// to the right, then jumping the item at k-to-the-right the next
// iteration.
func Cyclic(array []string, k int) (swaps int) {
	n := len(array)

	// This is the hard part: figuring out hom man times to jump
	// items, if you start at i in the array, and how many times
	// to start at progressive values of i
	hops := n
	max := gcd(n, k)
	if n%k == 0 {
		hops = n / k
		max = k
	}

	for i := 0; i < max; i++ {
		holder := array[i]
		for j, c := i, 0; c < hops; c++ {
			nxt := (j + k) % n
			tmp := array[nxt]
			array[nxt] = holder
			swaps++
			holder = tmp
			j = nxt
			if j == i {
				break
			}
		}
	}
	return
}

// Euclid's algorithm for finding Greatest Common Divisor
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Reversal rotates an array by k positions right
// via clever reversals.
func Reversal(array []string, k int) (swaps int) {
	l := len(array)
	swaps = reverse(array[:l-k])
	swaps += reverse(array[l-k:])
	swaps += reverse(array)
	return
}

func reverse(array []string) (swaps int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		swaps++
		array[i], array[j] = array[j], array[i]
	}
	return
}
