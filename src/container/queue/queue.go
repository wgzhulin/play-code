package queue

import (
	"github.com/play-code/src/basedata/ele"
)

type Queue interface {
	Enqueue(ele *ele.Ele)
	Dequeue() *ele.Ele
	IsEmpty() bool
	Full() bool
}
