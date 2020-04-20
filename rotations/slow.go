package rotations

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
