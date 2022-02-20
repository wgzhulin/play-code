package basicgo

import (
	"fmt"
	"testing"
)

type A struct {
	member int
}

func (a A) String() string {
	return fmt.Sprintf("%v", a.member)
}

type B struct {
	a      A
	aPoint *A
}

func (b B) String() string {
	return fmt.Sprintf("A: %v, aPoint: %v", b.a, b.aPoint)
}

// aPoint的地址是切片中对应的地址。切片中移除中间元素后会进行重排序
func TestSliceRemoveMidEle(t *testing.T) {
	demo := []B{
		{a: A{member: 1}},
		{a: A{member: 2}},
		{a: A{member: 3}},
	}

	for i := range demo {
		demo[i].aPoint = &demo[i].a
	}

	fmt.Println("demo before operation is ", demo) // result: demo before operation is  [A: 1, aPoint: 1 A: 2, aPoint:2 A: 3, aPoint: 3]
	demo = append(demo[:0], demo[1:]...)
	fmt.Println("demo after operation is ", demo) // result: demo after operation is  [A: 2, aPoint: 3 A: 3, aPoint: 3]
}

func TestAddSliceRemoveMidEle(t *testing.T) {
	demo := make([]B, 0, 3)
	a1 := A{member: 1}
	a2 := A{member: 2}
	a3 := A{member: 3}
	demo = append(demo, B{a: a1, aPoint: &a1}, B{a: a2, aPoint: &a2}, B{a: a3, aPoint: &a3})

	fmt.Println("demo before operation is ", demo) // result: demo before operation is  [A: 1, aPoint: 1 A: 2, aPoint: 2 A: 3, aPoint: 3]
	demo = append(demo[:0], demo[1:]...)
	fmt.Println("demo after operation is ", demo) // result: demo after operation is  [A: 2, aPoint: 2 A: 3, aPoint: 3]
}

