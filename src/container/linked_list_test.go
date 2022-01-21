package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkListPush(t *testing.T) {
	head := NewLinkedList()

	node1 := NewNode(1)
	head.pushBack(node1)
	assert.Equal(t, head.size(), 1)

	node2 := NewNode(2)
	head.pushBack(node2)
	assert.Equal(t, head.size(), 2)

	assert.True(t, head.node.equal(node1))
	assert.True(t, head.node.Next.equal(node2))
}

func TestLinkListPushBackFront(t *testing.T) {
	head := NewLinkedList()

	node1 := NewNode(1)
	head.pushBack(node1)
	assert.Equal(t, head.size(), 1)

	node2 := NewNode(2)
	head.pushFront(node2)
	assert.Equal(t, head.size(), 2)

	node3 := NewNode(3)
	head.pushFront(node3)
	assert.Equal(t, head.size(), 3)

	node4 := NewNode(4)
	head.pushBack(node4)
	assert.Equal(t, head.size(), 4)

	assert.True(t, head.node.equal(node3))
	assert.True(t, head.node.Next.equal(node2))
	assert.True(t, head.node.Next.Next.equal(node1))
	assert.True(t, head.node.Next.Next.Next.equal(node4))
}

func TestLinkListPop(t *testing.T) {
	head := NewLinkedList()

	node1 := NewNode(1)
	head.pushBack(node1)

	node2 := NewNode(2)
	head.pushBack(node2)

	back := head.popBack()
	assert.Equal(t, head.size(), 1)
	assert.True(t, back.equal(node2))

	front := head.popFront()
	assert.Equal(t, head.size(), 0)
	assert.True(t, front.equal(node1))
}

func TestLinkListBackFront(t *testing.T) {
	head := NewLinkedList()

	node1 := NewNode(1)
	head.pushBack(node1)

	node2 := NewNode(2)
	head.pushBack(node2)

	back := head.back()
	assert.Equal(t, head.size(), 2)
	assert.True(t, back.equal(node2))

	front := head.front()
	assert.Equal(t, head.size(), 2)
	assert.True(t, front.equal(node1))
}

func TestLinkListRemove(t *testing.T) {
	head := NewLinkedList()

	node1 := NewNode(1)
	head.pushBack(node1)
	assert.Equal(t, head.size(), 1)

	node2 := NewNode(2)
	head.pushFront(node2)
	assert.Equal(t, head.size(), 2)

	node3 := NewNode(3)
	head.pushFront(node3)
	assert.Equal(t, head.size(), 3)

	node4 := NewNode(4)
	head.pushBack(node4)
	assert.Equal(t, head.size(), 4)

	head.removeFirstValue(node3)
	assert.True(t, head.node.equal(node2))
	assert.True(t, head.node.Next.equal(node1))
	assert.True(t, head.node.Next.Next.equal(node4))

	head.remove(1)
	assert.True(t, head.node.equal(node2))
	assert.True(t, head.node.Next.equal(node4))
}
