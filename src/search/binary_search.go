package search

import "math"

func BinarySearch(s []int, ele int) int {
	low := 0
	high := len(s) - 1
	for low <= high {
		mid := low + (high-low)/2
		midValue := s[mid]
		if ele == midValue {
			return ele
		}
		if ele < midValue {
			high = mid - 1
		}
		if ele > midValue {
			low = mid + 1
		}
	}
	return math.MaxInt32 // not find
}
