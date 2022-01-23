package container

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueueFixArray(t *testing.T) {
	q := NewQueueFixArray()

	input := make([]*Ele, 10)
	for i := 0; i < 10; i++ {
		input[i] = NewEle(i + 1)
	}

	for i := range input {
		q.enqueue(input[i])
	}
	assert.True(t, q.full())

	for i := range input {
		assert.True(t, q.dequeue().Equals(input[i]))
	}

}

func TestNewQueueFixArrayHalf(t *testing.T) {
	q := NewQueueFixArray()

	input := make([]*Ele, 5)
	for i := 0; i < 5; i++ {
		input[i] = NewEle(i + 1)
	}

	for i := range input {
		q.enqueue(input[i])
	}

	ele1 := q.dequeue()
	assert.True(t, ele1.Equals(input[0]))

	secondInput := make([]*Ele, 6)
	for i := 0; i < 6; i++ {
		secondInput[i] = NewEle(i + 6)
	}

	for i := range secondInput {
		q.enqueue(secondInput[i])
	}

	allInputs := append(input[1:], secondInput...)

	for i := range allInputs {
		assert.True(t, q.dequeue().Equals(allInputs[i]))
	}
}

func TestNewQueueFixArrayEnqueueDequeue(t *testing.T) {
	q := NewQueueFixArray()

	input := make([]*Ele, 10)
	for i := 0; i < 10; i++ {
		input[i] = NewEle(i + 1)
	}

	for i := range input {
		q.enqueue(input[i])
		assert.True(t, q.dequeue().Equals(input[i]))
	}
	assert.True(t, q.isEmpty())
	assert.Equal(t, q.size, 0)
	assert.Equal(t, q.cur, queueMaxSize-1)
}

func TestNewEle(t *testing.T) {
	ele1 := NewEle(1)
	ele := NewEle(1)
	fmt.Println(ele.Equals(ele1))
}
