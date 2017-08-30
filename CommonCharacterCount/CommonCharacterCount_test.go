package commonCharacterCount

import (
	"log"
	"testing"
	"time"
)

type testPair struct {
	s1       string
	s2       string
	expected int
}

var testInputs = []testPair{
	{"aabcc", "adcaa", 3},
	{"adcaa", "aabcc", 3},
	{"", "aaacc", 0},
	{"zzz", "ccc", 0},
}

func TestCommonCharacterCount(t *testing.T) {
	start := time.Now()
	for _, input := range testInputs {
		result := commonCharacterCount(input.s1, input.s2)
		if input.expected != result {
			t.Error(
				"For", input.s1, input.s2,
				"expected", input.expected,
				"got", result,
			)
		}
		t.Log(
			"For", input.s1, input.s2,
			"expected", input.expected,
			"got", result,
		)
	}
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

func TestCommonCharacterCount2(t *testing.T) {
	start := time.Now()
	for _, input := range testInputs {
		result := commonCharacterCount2(input.s1, input.s2)
		if input.expected != result {
			t.Error(
				"For", input.s1, input.s2,
				"expected", input.expected,
				"got", result,
			)
		}
		t.Log(
			"For", input.s1, input.s2,
			"expected", input.expected,
			"got", result,
		)
	}
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}
