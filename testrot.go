package main

import (
	"fmt"
	"reflect"

	"array-rotation-algorithms/rotations"
)

type RotFunc func([]string, int) int
type RotateFunc struct {
	fn   func([]string, int) int
	name string
}

var rotationfunctions = []RotateFunc{
	{fn: rotations.Cyclic, name: "Cyclic  "},
	{fn: rotations.Reversal, name: "Reversal"},
	{fn: rotations.Rotate, name: "Rotate  "},
}

func main() {

	orig := []string{
		`a`, `b`, `c`, `d`, `e`, `f`, `g`, `h`, `i`, `j`, `k`, `l`, `m`,
		`n`, `o`, `p`, `q`, `r`, `s`, `t`, `u`, `v`, `w`, `x`, `y`, `z`,
		`A`, `B`, `C`, `D`, `E`, `F`, `G`, `H`, `I`, `J`, `K`, `L`, `M`,
		`N`, `O`, `P`, `Q`, `R`, `S`, `T`, `U`, `V`, `W`, `X`, `Y`, `Z`,
		`1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `0`, `~`, `!`, `@`,
		`#`, `$`, `%`, `^`, `&`, `*`, `(`, `)`, `_`, `-`, `+`, `=`, `{`,
		`[`, `}`, `]`, `|`, `:`, `;`, `"`, `'`, `<`, `,`, `>`, `.`, `?`,
		`/`,
	}

	for L := 2; L < len(orig); L++ {
		for N := 1; N < L; N++ {
			std := make([]string, L)
			copy(std, orig)
			rotations.Slowrotate(std, N)
			// std is the correct rotation, I hope

			for _, rfn := range rotationfunctions {
				array := make([]string, L)
				copy(array, orig)
				rfn.fn(array, N)

				if reflect.DeepEqual(std, array) {
					fmt.Printf("%s worked %2d -> %2d\n", rfn.name, L, N)
				} else {
					fmt.Printf("%s failed %2d -> %2d\n", rfn.name, L, N)
					fmt.Printf("%v\n", std)
					fmt.Printf("%v\n", array)
					return
				}
			}
		}
	}

}
