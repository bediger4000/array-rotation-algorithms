package rotations

import "fmt"

// Cyclic rotates an array k items to the right by jumping each item k
// to the right, then jumping the item at k-to-the-right the next
// iteration. The big problem with this is figuring out when to
// stop a cylce of jumps.
func Cyclic(array []string, k int) (swaps int) {
	n := len(array)

	// This is the hard part: figuring out hom many times to jump
	// items if you start at i in the array, and how many times
	// to start at successive values of i

	hops := n
	max := gcd(n, k)

	// n a multiple of k: incrementing some index i by k
	// enough times and you'll get back to i, because we're
	// doing (i + m*k)%n to wrap the index around the array end
	if n%k == 0 {
		hops = n / k
		max = k
	}
	fmt.Printf("Array length n %d, k %d, hops %d, max %d\n",
		n, k, hops, max)

	for i := 0; i < max; i++ {
		holder := array[i]
		for j, c := i, 0; c < hops; c++ {
			fmt.Printf("j %d, c %d\n", j, c)
			nxt := (j + k) % n
			tmp := array[nxt]
			array[nxt] = holder
			swaps++
			holder = tmp
			j = nxt
			if j == i {
				fmt.Printf("j == %d, break\n", i)
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
