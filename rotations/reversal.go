package rotations

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
