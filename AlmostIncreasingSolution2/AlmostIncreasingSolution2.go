package solution2

func almostIncreasingSequence(sequence []int) bool {
	// no error
	err := 0
	// [ 0, 1, 2, 3, 2, 3, 5 ]
	// starting from 1 since we compare s[i] with s[i -1]
	for i := 1; i < len(sequence); i++ {

		// if current is less than previous, err occured
		if sequence[i] <= sequence[i-1] {
			//err count +1
			err++
			// if this is second err, proven not valid sequence
			if err > 1 {
				return false
			}

			// i+1 < len(sequence) protects us from exceeding upper bound of i
			// i-2 >= 0 protects us from exceeding lower bound of i
			// next is detecting absolute invalid sequence.
			//
			// When we encounter an error, it means either current or previous int is wrong
			// We can proceed to the next check as long as one of them is viable
			// case1: sequence[i-2] >= sequence[i], sequence[i] is wrong
			// case2: sequence[i+1] <= sequence[i-1], sequence[i-1] is wrong
			// if both cases are true, we have a invalid sequence
			// example: [  1,  4,  5,  3,  4,  6,  7]
			//                 ^   ^   ^   ^
			//                 |   ----|----
			//                 ---------
			if i+1 < len(sequence) && i-2 >= 0 && sequence[i-2] >= sequence[i] && sequence[i+1] <= sequence[i-1] {
				return false
			}
		}
	}
	return true
}
