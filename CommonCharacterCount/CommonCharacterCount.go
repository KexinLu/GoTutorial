package commonCharacterCount

import (
	"sort"
)

type runeSlice []rune

func (rs runeSlice) Len() int {
	return len(rs)
}

func (rs runeSlice) Less(i int, j int) bool {
	return rs[i] < rs[j]
}

func (rs runeSlice) Swap(i int, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func commonCharacterCount(s1 string, s2 string) int {
	common := 0
	r1 := runeSlice(s1)
	r2 := runeSlice(s2)
	sort.Sort(r1)
	sort.Sort(r2)
	p1 := 0
	p2 := 0
	for p1 < len(r1) && p2 < len(r2) {
		if r1[p1] == r2[p2] {
			common++
			p1++
			p2++
		} else {
			if int32(r1[p1]) > int32(r2[p2]) {
				p2++
			} else {
				p1++
			}
		}
	}

	return common
}

func commonCharacterCount2(s1 string, s2 string) int {

	sum := 0
	m := make(map[rune]int)
	for _, r := range s1 {
		m[r]++
	}
	for _, r := range s2 {
		if m[r] > 0 {
			sum++
			m[r]--
		}
	}
	return sum
}
