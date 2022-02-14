package queue

import (
	"github.com/play-code/src/basic/datastructure/basedata/ele"
	"github.com/play-code/src/basic/datastructure/basedata/node"
	"github.com/play-code/src/basic/datastructure/container/list"
)

type LinkedListQueue struct {
	data list.LinkedList
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (q *LinkedListQueue) Enqueue(ele *ele.Ele) {
	if q == nil || q.Full() {
		return
	}
	q.data.PushBack(node.NewNode(ele))
}

func (q *LinkedListQueue) Dequeue() *ele.Ele {
	if q.IsEmpty() {
		return nil
	}

	if frontNode := q.data.PopFront(); frontNode != nil {
		if value, ok := frontNode.Value.(*ele.Ele); ok {
			return value
		}
	}

	return nil
}

func (q *LinkedListQueue) IsEmpty() bool {
	if q == nil {
		return true
	}
	return q.data.Size() == 0
}

func (q *LinkedListQueue) Full() bool {
	if q == nil {
		return true
	}
	return q.data.Size() == queueMaxSize
}
