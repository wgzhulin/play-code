package list

import (
	"github.com/play-code/src/basic/datastructure/basedata/node"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkListPush(t *testing.T) {
	head := NewLinkedList()

	node1 := node.NewNode(1)
	head.PushBack(node1)
	assert.Equal(t, head.Size(), 1)

	node2 := node.NewNode(2)
	head.PushBack(node2)
	assert.Equal(t, head.Size(), 2)

	assert.True(t, head.node.Equals(node1))
	assert.True(t, head.node.Next.Equals(node2))
}

func TestLinkListPushBackFront(t *testing.T) {
	head := NewLinkedList()

	node1 := node.NewNode(1)
	head.PushBack(node1)
	assert.Equal(t, head.Size(), 1)

	node2 := node.NewNode(2)
	head.PushFront(node2)
	assert.Equal(t, head.Size(), 2)

	node3 := node.NewNode(3)
	head.PushFront(node3)
	assert.Equal(t, head.Size(), 3)

	node4 := node.NewNode(4)
	head.PushBack(node4)
	assert.Equal(t, head.Size(), 4)

	assert.True(t, head.node.Equals(node3))
	assert.True(t, head.node.Next.Equals(node2))
	assert.True(t, head.node.Next.Next.Equals(node1))
	assert.True(t, head.node.Next.Next.Next.Equals(node4))
}

func TestLinkListPop(t *testing.T) {
	head := NewLinkedList()

	node1 := node.NewNode(1)
	head.PushBack(node1)

	node2 := node.NewNode(2)
	head.PushBack(node2)

	back := head.PopBack()
	assert.Equal(t, head.Size(), 1)
	assert.True(t, back.Equals(node2))

	front := head.PopFront()
	assert.Equal(t, head.Size(), 0)
	assert.True(t, front.Equals(node1))
}

func TestLinkListBackFront(t *testing.T) {
	head := NewLinkedList()

	node1 := node.NewNode(1)
	head.PushBack(node1)

	node2 := node.NewNode(2)
	head.PushBack(node2)

	back := head.Back()
	assert.Equal(t, head.Size(), 2)
	assert.True(t, back.Equals(node2))

	front := head.Front()
	assert.Equal(t, head.Size(), 2)
	assert.True(t, front.Equals(node1))
}

func TestLinkListRemove(t *testing.T) {
	head := NewLinkedList()

	node1 := node.NewNode(1)
	head.PushBack(node1)
	assert.Equal(t, head.Size(), 1)

	node2 := node.NewNode(2)
	head.PushFront(node2)
	assert.Equal(t, head.Size(), 2)

	node3 := node.NewNode(3)
	head.PushFront(node3)
	assert.Equal(t, head.Size(), 3)

	node4 := node.NewNode(4)
	head.PushBack(node4)
	assert.Equal(t, head.Size(), 4)

	head.RemoveFirstValue(node3)
	assert.True(t, head.node.Equals(node2))
	assert.True(t, head.node.Next.Equals(node1))
	assert.True(t, head.node.Next.Next.Equals(node4))

	head.Remove(1)
	assert.True(t, head.node.Equals(node2))
	assert.True(t, head.node.Next.Equals(node4))
}
