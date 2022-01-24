package search

import "math"

func RecurBinarySearch(s []int, ele int) int {
	return recurBs(s, 0, len(s)-1, ele)
}

// 递归
func recurBs(s []int, low, high, ele int) int {
	mid := low + (high-low)/2
	if low > high {
		return math.MaxInt32 // not find
	}

	if ele < s[mid] {
		return recurBs(s, low, mid-1, ele)
	} else if ele > s[mid] {
		return recurBs(s, mid+1, high, ele)
	} else {
		return ele
	}
}
