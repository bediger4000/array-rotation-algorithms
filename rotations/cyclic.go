package rotations

// Cyclic rotates an array k items to the right by jumping each item k
// to the right, then jumping the item at k-to-the-right the next
// iteration. The big problem with this is figuring out when to
// stop a cylce of jumps.
func Cyclic(array []string, k int) (swaps int) {
	n := len(array)

	// If n and k have a GCD of 1, you end up skipping through
	// all the indexes, next i = (i + m*k)%n.
	// If n and k have a GCD larger than 1, you skip through
	// the GCD number of sets of indexes. In any case, you have
	// to stop jumping if the next index you'd change is the index
	// you started at.

	max := gcd(n, k)

	for i := 0; i < max; i++ {
		holder := array[i]
		for j := i; true; {
			nxt := (j + k) % n
			tmp := array[nxt]
			array[nxt] = holder
			holder = tmp
			swaps++
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
