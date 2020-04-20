package main

import (
	"fmt"
	"rotate/rotations"
)

func main() {

	orig := []string{"a", "b", "c", "d", "e", "f", "g", "h",
		"i", "j", "k", "l", "m", "n", "o", "p"}

	for L := 2; L < len(orig); L++ {
		for N := 1; N < L; N++ {
			array1 := make([]string, L)
			array2 := make([]string, L)
			array3 := make([]string, L)
			array4 := make([]string, L)
			copy(array1, orig)
			copy(array2, orig)
			copy(array3, orig)
			copy(array4, orig)
			slow := rotations.Slowrotate(array1, N)
			cyclic := rotations.Cyclic(array2, N)
			revs := rotations.Reversal(array3, N)
			mine := rotations.Rotate(array4, N)
			fmt.Printf("%2d %2d  %3d %3d %3d %3d\n", L, N, slow, cyclic, revs, mine)
		}
	}

}
