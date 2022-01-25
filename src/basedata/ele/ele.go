package ele

import (
	"reflect"
)

type Ele struct {
	value interface{}
}

func NewEle(value interface{}) *Ele {
	return &Ele{value: value}
}

func (e *Ele) Equals(ele *Ele) bool {
	if e == nil || ele == nil {
		return false
	}
	return reflect.DeepEqual(e.value, ele.value)
}
