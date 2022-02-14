package basicgo

import "sort"

type SliceOper struct{}

func (i *SliceOper) Reverse(s []int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		j := len(s) - 1 - i
		s[i], s[j] = s[j], s[i]
	}
}

func (i *SliceOper) Sort(s []int) {
	sort.Ints(s)
}

func (i *SliceOper) RemoveDeduplicateOrderSlice(s []int) []int {
	i.Sort(s)
	j := 0
	for i := 1; i < len(s); i++ {
		if s[j] == s[i] {
			continue
		}
		j++
		s[j] = s[i]
	}
	return s[:j+1]
}

func (i *SliceOper) SlidingWindow(s []int, wdSize int) [][]int {
	if len(s) <= wdSize {
		return [][]int{s}
	}
	r := make([][]int, 0, len(s)-wdSize+1)
	for i, j := 0, wdSize; j <= len(s); i, j = i+1, j+1 {
		r = append(r, s[i:j])
	}

	return r
}
