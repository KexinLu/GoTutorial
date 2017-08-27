package almostincreasing

import (
	"testing"
)

type testPair struct {
	intSlice []int
	expected bool
}

var testInputs = []testPair{
	{[]int{1, 2, 3, 4}, true},
	{[]int{1, 2, 2, 4}, true},
	{[]int{3, 2, 2, 4}, false},
	{[]int{10, 1, 2, 3, 4}, true},
	{[]int{0, 1, 2, 2, 2, 5}, false},
}

func TestAlmostIncreasingSequence(t *testing.T) {
	for _, input := range testInputs {
		result := almostIncreasingSequence(input.intSlice)
		if input.expected != result {
			t.Error(
				"For", input.intSlice,
				"expected", input.expected,
				"got", result,
			)
		}
		t.Log(
			"For", input.intSlice,
			"expected", input.expected,
			"got", result,
		)
	}
}
