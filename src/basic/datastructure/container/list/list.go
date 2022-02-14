package list

type List []interface{}

func NewList() List {
	return make(List, 0)
}

func (l *List) Add(ele interface{}) {
	*l = append(*l, ele)
}

func (l *List) Clear() {
	if len(*l) == 0 {
		return
	}
	for i := range *l {
		(*l)[i] = nil
	}
	*l = (*l)[:0]
}

func (l *List) Remove(index int) interface{} {
	if len(*l) == 0 || index <0 || index > len(*l)-1{
		return nil
	}
	result := (*l)[index]
	*l = append((*l)[:index], (*l)[index+1:]...)
	(*l)[len(*l)-1] = nil

	return result
}

func (l *List) Size() int {
	return len(*l)
}

func (l *List) IsEmpty() bool {
	return len(*l) == 0
}

func (l *List) Contains(ele interface{}) bool {
	if len(*l) == 0 {
		return false
	}

	for i := range *l {
		if (*l)[i] == ele {
			return true
		}
	}
	return false
}