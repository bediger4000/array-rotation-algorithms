# Daily Coding Problem: Problem #197 [Easy]

I got a "programming puzzle of the day [easy]" that said this:

	Given an array and a number k that's smaller than the length of the
    array, rotate the array to the right k elements in-place.

The tricky thing is the "in place" requirement,
otherwise the easy, obvious thing is to copy elements into
a separate array in rotation-sequence,
using some clever indexing involving `(i + k)%n` where `n` is the length of the array.

There's an easy *O(n\*k)* solution, rotate the array right by 1, k times.
Here it is in [Go](https://golang.org/):

```go
// Rotate an array by 1 rotate number of times
for i := 0; i < k; i++ {
    // Rotate an array by 1
    tmp := array[i]
    for j := i + 1; j != i; j = (j + 1) % len(array) {
        array[j], tmp = tmp, array[j]
    }
    array[i] = tmp
}
```

It does require space for 1 extra item
which is hidden by Golang multiple assignments,
but you could do that dumb swap-by-XOR stunt to avoid that.

As near as I can tell,
there's two other commonly given ways of doing this.

1. [cyclic replacements](https://leetcode.com/articles/rotate-array/#)
2. [repeated reversals](https://www.geeksforgeeks.org/program-for-array-rotation-continued-reversal-algorithm/)

Like a lot of this sort of programming problem,
the search space is polluted by vaguely correct explanations
and implementations, written mainly for absolute beginners.
The above examples are pretty weak,
I've re-implemented them in Go.

## My Solution

This discussion is for an array of length n, 0-indexed,
rotated k spots to the right.

My solution uses the first k spots in the array as temporary
space, swapping from `array[0]` with `array[m*k]`,
`array[1]` with `array[m*k + 1]`,
`array[2]` with `array[m*k + 2]`,
up to 
`array[k-1]` with `array[m*k + (k-1)]`.
The stride counter `m` gets incremented every `k` swaps.
The code uses two for-loops to do this.
On each swap, update a variable with the index `m*k + (k-1)`.
This values gets used to decide a final rotation later.

Do this in blocks of k values, up to the end of the array.
Because I wrote this in Go,
and I'm using slices,
this code technically does rotate the array in place,
or at least with only a few extra words of memory.

If k is a divisor of n (length of array),
the last block of k items got moved to indxes 0:k-1,
the algorithm has finished.

This is the best case, if we're concerned about number of swaps.
For instance, rotating an array of length 9 by 3:

1. Values at indexes 0, 1, 2 swapped with values at indexes 3, 4, 5
2. Values at indexes 0, 1, 2 swapped with values at indexes 6, 7, 8

That puts everything at the right place with only 6 swaps,
because the 3 temporary slots are the 3 final slots, too.

if k is not a divsor of n, there are two cases:

1. The algorithm has not swapped 1 or more items.
This can happen when you rotate by more than half the length of the array,
like rotating a length 7 array by 4.
The algorithm will move items indexed 0,1,2 to 4,5,6,
leaving the item at index 3 untouched.
2. The algorithm has swapped every item,
but less than k items have swapped into their final position.

Case 1, one or more items haven't been swapped,
either to the temporary area, or to their final destination.
But the temporary area plus the unmoved items can be rotated
by some number to get them,
and the rest of the temporary area valuse,
into their final positions.

Consider an array `[a, b, c, d, e, f, g, h, i, j]` of length 10,
to be rotated 7 to the right.
After the moving values a, b, c to their final positions,
the array (doing this in-place) looks like this:

    [h, i, j, d, e, f, g, a, b, c]

Items 'd', 'e' haven't moved, items 'h', 'i', 'j', are in the temp area.
If we rotate the subarray `[h, i, j, d, e]` by the number of untouched
values, we end up with a correctly rotated array.
Rotating the subarray of elements swapped,
but not in their final spots, and the unmoved elements works because
the first set of moves always ends putting an element in the last
position of the array.
The unmoved elements new indexes "wrap around" to indexes of
elements in the temp area, currently occupied by values not in their
ultimate correct position.

Case 2, all items have been swapped to the temporary sub-array.
The items might not have been swapped into the correct position
in the temporary area.
Since the algorithm uses the temporary area indexes in order,
all the values in it retain the correct rotational positions,
and the temporary area can be rotated into the correct position.
By keeping track of what index in the temporary area
(which is k elements long) the last value got placed in,
we can calculate by how many to rotate the temp area.

## Programs

I chose to have `[]string` as the array,
so that I could see what rotations (or failures) took place
without confusing numeric values in arrays and their numeric indexes.

I don't think there needs to be a makefile to codify builds.
`go build testrot.go` or whatever should get things to build.
I did this on Linux, with Go `go1.14.2 linux/amd64` compiler,
but there's nothing outrageously un-portable here.

The actual rotation algorithms are in `package rotations`,
one file per algorithm.

* `testrot.go` - systematically try all rotations of n > 1 length arrays,
and all values of k (the rotation) that make sense (1, k-1].
Stops on the first failure.
* `swapcounter.go` - counts the number of swaps of values between array indexes
for all the array rotating algorithms.
* `rev.go` - program to exercize one method of array rotation more or less
interactively. To see an array of length 5 rotated 3, run: `./rev 3 a b c d e`
