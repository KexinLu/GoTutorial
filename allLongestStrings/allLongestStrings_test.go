package allLongStrings

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type testPair struct {
	stringSlice []string
	expected    []string
}

var inputs = []testPair{
	{[]string{"12", "abc", "cccc", "dddd"}, []string{"cccc", "dddd"}},
	{[]string{"12", "abc", "aaa", "zzz"}, []string{"aaa", "zzz", "abc"}},
	{[]string{"1", "2", "c", "d"}, []string{"1", "2", "c", "d"}},
}

func TestAllLongString(t *testing.T) {
	for _, input := range inputs {
		result := allLongStrings(input.stringSlice)

		// DeepEqual is slow, but for our purpose it's ok. In real project
		// we need to implement our own comparasion
		if !sliceEqual(input.expected, result) {
			t.Error(
				"For", input.stringSlice,
				"expected", input.expected,
				"got", result,
			)
		}
		t.Log(
			"For", input.stringSlice,
			"expected", input.expected,
			"got", result,
		)
	}
}

// this is a hack
func sliceEqual(s1 []string, s2 []string) bool {
	sort.Strings(s1)
	sort.Strings(s2)
	fmt.Printf("%v", reflect.DeepEqual(s1, s2))
	return reflect.DeepEqual(s1, s2)
}
