package queue

import (
	"github.com/play-code/src/basic/datastructure/basedata/ele"
)

// max size of the queue
const queueMaxSize = 10

type FixArrayQueue struct {
	size int
	cur  int
	data FixArray
}

type FixArray [queueMaxSize]*ele.Ele

func NewFixArrayQueue() *FixArrayQueue {
	return &FixArrayQueue{cur: queueMaxSize - 1}
}

func (q *FixArrayQueue) Enqueue(ele *ele.Ele) {
	if q == nil || q.Full() {
		return
	}
	cur := (q.cur - q.size + queueMaxSize) % queueMaxSize
	q.data[cur] = ele
	q.size++
}

func (q *FixArrayQueue) Dequeue() *ele.Ele {
	if q.IsEmpty() {
		return nil
	}

	result := q.data[q.cur]
	q.data[q.cur] = nil
	q.cur--
	if q.cur < 0 {
		q.cur = queueMaxSize - 1
	}
	q.size--
	return result
}

func (q *FixArrayQueue) IsEmpty() bool {
	if q == nil {
		return true
	}
	return q.size == 0
}

func (q *FixArrayQueue) Full() bool {
	if q == nil {
		return true
	}
	return q.size == queueMaxSize
}
