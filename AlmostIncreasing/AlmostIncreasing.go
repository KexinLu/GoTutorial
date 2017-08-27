package almostincreasing

func almostIncreasingSequence(sequence []int) bool {
	// First look at isIncreasingSequence
	errPos, increasing := isIncreasingSequence(sequence)
	// if increasing is true, then no error, return true
	if increasing {
		return true
	} else {
		// increasing is not true, this means we need to see if we can remove one and make it work
		// two possible way, remove last, or remove current
		// example:
		// [10, 1, 2, 3, 4]      errPos => 0,  we should remove element errPos = 0 (10) and it will work
		possible1 := takeOutElementAtPosition(sequence, errPos)
		// [4, 5, 3, 6, 10]      errPos => 1,  we should remove element errPos + 1 = 2 (3) and it will work
		possible2 := takeOutElementAtPosition(sequence, errPos+1)
		// uncomment to see the values of the two slices
		//fmt.Printf("%v", possible1)
		//fmt.Printf("%v", possible2)

		// since this is double check, we don't care about error position, throw'em out with _
		_, incr1 := isIncreasingSequence(possible1)
		_, incr2 := isIncreasingSequence(possible2)
		// either one true indicates a viable slice
		return incr1 || incr2
	}
}

func takeOutElementAtPosition(sequence []int, position int) []int {
	// make a copy of origional slice.
	// this is not optimal technically we only need to copy once, but for the sake of reader make it easier
	// reason for making a copy is that append will modify the original slice.
	copyOfSec := make([]int, len(sequence)) // make it with same length
	copy(copyOfSec, sequence)               // copy from sequence to copyOfSec
	//                 [ *, *, *, *, err, * ,* , *]
	//     append      [           ],     [       ]
	//                       ^                 ^
	return append(copyOfSec[:position], copyOfSec[position+1:]...)
}

// input is slice  output is an int and a boolean
func isIncreasingSequence(sequence []int) (int, bool) {
	last := sequence[0] // first element in slice
	current := 0        // current Element, set to 0 as default value.  same as   var current int

	// skip i = 0 since last is sequence[0], starting from 1
	// example:
	// slice:     [    3,     5,     9,    11,    2,    19]
	// position:       0      1      2     3      4       5
	//               last  current                          // i = 1
	//                      last  current                   // i = 2
	//                            last  current             // i = 3
	//                                   last  current      // i = 3 found err, since current <= last
	for i := 1; i < len(sequence); i++ {
		current = sequence[i]
		if current <= last {
			// found error, i - 1 is last's position, false means this is not increasing sequence
			return i - 1, false
		}
		// did not find error, replace last with current
		last = current
	}
	// after looping through enitre slice, no error found
	// error position is -1 (will be ignored)
	return -1, true
}
