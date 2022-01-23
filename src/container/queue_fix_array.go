package container

// max size of the queue
const queueMaxSize = 10

type QueueFixArray struct {
	size int
	cur  int
	data FixArray
}

type FixArray [queueMaxSize]*Ele

func NewQueueFixArray() *QueueFixArray {
	return &QueueFixArray{cur: queueMaxSize - 1}
}

func (q *QueueFixArray) enqueue(ele *Ele) {
	if q == nil || q.full() {
		return
	}
	cur := (q.cur - q.size + queueMaxSize) % queueMaxSize
	q.data[cur] = ele
	q.size++
}

func (q *QueueFixArray) dequeue() *Ele {
	if q.isEmpty() {
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

func (q *QueueFixArray) isEmpty() bool {
	if q == nil {
		return true
	}
	return q.size == 0
}

func (q *QueueFixArray) full() bool {
	if q == nil {
		return true
	}
	return q.size == queueMaxSize
}
