package queue

import (
	"github.com/play-code/src/basic/datastructure/basedata/ele"
)

type Queue interface {
	Enqueue(ele *ele.Ele)
	Dequeue() *ele.Ele
	IsEmpty() bool
	Full() bool
}

func NewQueue() Queue {
	return NewLinkedListQueue()
}
