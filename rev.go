package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"array-rotation-algorithms/rotations"
)

func main() {

	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var orig []string

	for _, str := range os.Args[2:] {
		orig = append(orig, str)
	}

	n := len(orig)
	cpy := make([]string, n)
	copy(cpy, orig)
	rotations.Slowrotate(cpy, k)

	array1 := make([]string, n)
	copy(array1, orig)
	swaps := rotations.Reversal(array1, k)

	if reflect.DeepEqual(cpy, array1) {
		fmt.Printf("Reversal worked %2d -> %2d, %d swaps\n", n, k, swaps)
	} else {
		fmt.Printf("Reversal failed %2d -> %2d\n", n, k)
		fmt.Printf("%v\n", cpy)
		fmt.Printf("%v\n", array1)
	}

}
