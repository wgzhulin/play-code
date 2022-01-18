package container

type List []interface{}

func NewList() List {
	return make(List, 0)
}

func (l *List) add(ele interface{}) {
	*l = append(*l, ele)
}

func (l *List) clear() {
	if len(*l) == 0 {
		return
	}
	for i := range *l {
		(*l)[i] = nil
	}
	*l = (*l)[:0]
}

func (l *List) remove(index int) interface{} {
	if len(*l) == 0 || index <0 || index > len(*l)-1{
		return nil
	}
	result := (*l)[index]
	*l = append((*l)[:index], (*l)[index+1:]...)
	(*l)[len(*l)-1] = nil

	return result
}

func (l *List) size() int {
	return len(*l)
}

func (l *List) isEmpty() bool {
	return len(*l) == 0
}

func (l *List) contains(ele interface{}) bool {
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