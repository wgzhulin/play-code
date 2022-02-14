package basicgo

import (
	"fmt"
	"testing"
)

func TestSliceIntCopy(t *testing.T) {
	a := []int{1}
	a = append(a, 2)
	b := make([]int, len(a))
	copy(b, a) // return the number of the element copied

	fmt.Printf("aPoint: %p, A: %v\n", a, a)
	fmt.Printf("bPoint: %p, b: %v\n", b, b)
}

func TestSliceExpandCopy(t *testing.T) {
	a := []int{1}
	a = append(a, 2)

	b := append(a[0:0:0], a...)

	fmt.Printf("aPoint: %p, A: %v\n", a, a)
	fmt.Printf("bPoint: %p, b: %v\n", b, b)
}

func TestSliceNilIntCopy(t *testing.T) {
	a := []int{1}
	a = append(a, 2)

	b := append([]int(nil), a...)

	fmt.Printf("aPoint: %p, A: %v\n", a, a)
	fmt.Printf("bPoint: %p, b: %v\n", b, b)
}

// A slice store *A
func TestSlicePointCopy(t *testing.T) {
	type A struct {
		value int
	}
	printer := func(s []*A) string {
		result := ""
		for _, v := range s {
			result += fmt.Sprintf("%v ", v.value)
		}
		return result
	}

	a1 := A{value: 1}
	a2 := A{value: 2}
	aSlice := []*A{&a1}
	aSlice = append(aSlice, &a2)
	bSlice := make([]*A, len(aSlice))
	copy(bSlice, aSlice) // return the number of the element copied

	fmt.Printf("aSlicePoint: %p, aSlice: %v, aSliceValue: %v\n", aSlice, aSlice, printer(aSlice))
	fmt.Printf("bSlicePoint: %p, bSlice: %v, aSliceValue: %v\n", bSlice, bSlice, printer(bSlice))

	// change slice A value will also changer slice b value
	// so copy func is only copy the value of slice
	aSlice[0].value = 3
	fmt.Printf("aSlicePoint: %p, aSlice: %v, aSliceValue: %v\n", aSlice, aSlice, printer(aSlice))
	fmt.Printf("bSlicePoint: %p, bSlice: %v, aSliceValue: %v\n", bSlice, bSlice, printer(bSlice))

}