package queue

import (
	"github.com/play-code/src/basic/datastructure/basedata/ele"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLinkedListArray(t *testing.T) {
	q := NewLinkedListQueue()

	input := make([]*ele.Ele, 10)
	for i := 0; i < 10; i++ {
		input[i] = ele.NewEle(i + 1)
	}

	for i := range input {
		q.Enqueue(input[i])
	}
	assert.True(t, q.Full())

	for i := range input {
		assert.Equal(t, q.Dequeue(), input[i])
	}

}

func TestNewLinkedListArrayHalf(t *testing.T) {
	q := NewLinkedListQueue()

	input := make([]*ele.Ele, 5)
	for i := 0; i < 5; i++ {
		input[i] = ele.NewEle(i + 1)
	}

	for i := range input {
		q.Enqueue(input[i])
	}

	ele1 := q.Dequeue()
	assert.True(t, ele1.Equals(input[0]))

	secondInput := make([]*ele.Ele, 6)
	for i := 0; i < 6; i++ {
		secondInput[i] = ele.NewEle(i + 6)
	}

	for i := range secondInput {
		q.Enqueue(secondInput[i])
	}

	allInputs := append(input[1:], secondInput...)

	for i := range allInputs {
		assert.True(t, q.Dequeue().Equals(allInputs[i]))
	}
}

func TestNewLinkedListArrayEnqueueDequeue(t *testing.T) {
	q := NewLinkedListQueue()

	input := make([]*ele.Ele, 10)
	for i := 0; i < 10; i++ {
		input[i] = ele.NewEle(i + 1)
	}

	for i := range input {
		q.Enqueue(input[i])
		assert.Equal(t, q.data.Size(), 1)
		assert.True(t, q.Dequeue().Equals(input[i]))
		assert.Equal(t, q.data.Size(), 0)
	}

	assert.True(t, q.IsEmpty())
}
