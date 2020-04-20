package main

import (
	"fmt"
	"reflect"

	"rotate/rotations"
)

func main() {

	orig := []string{"a", "b", "c", "d", "e", "f", "g", "h",
		"i", "j", "k", "l", "m", "n", "o", "p"}

	for L := 2; L < len(orig); L++ {
		for N := 1; N < L; N++ {
			cpy := make([]string, L)
			copy(cpy, orig)
			rotations.Slowrotate(cpy, N)

			array1 := make([]string, L)
			copy(array1, orig)
			rotations.Rotate(array1, N)

			if reflect.DeepEqual(cpy, array1) {
				fmt.Printf("Rotate worked %2d -> %2d\n", L, N)
			} else {
				fmt.Printf("Rotate failed %2d -> %2d\n", L, N)
				fmt.Printf("%v\n", cpy)
				fmt.Printf("%v\n", array1)
			}

			array2 := make([]string, L)
			copy(array2, orig)
			rotations.Cyclic(array2, N)

			if reflect.DeepEqual(cpy, array2) {
				fmt.Printf("Cyclic rotation worked %2d -> %2d\n", L, N)
			} else {
				fmt.Printf("Cyclic rotation failed %2d -> %2d\n", L, N)
				fmt.Printf("%v\n", cpy)
				fmt.Printf("%v\n", array2)
			}
		}
	}

}
