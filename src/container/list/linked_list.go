package list

import (
	node2 "github.com/play-code/src/basedata/node"
)

// Header
type LinkedList struct {
	num  int
	node *node2.Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Size() int {
	if l.IsEmpty() {
		return 0
	}
	return l.num
}

func (l *LinkedList) Insert(index int, node *node2.Node) {
	if l.IsEmpty() || index > l.num-1 {
		return
	}
	if index == 0 {
		l.PushFront(node)
	} else {
		preNode := l.ValueAtIndex(index - 1)
		node.Next = preNode.Next.Next
		preNode.Next = node
	}
	l.num++
}

// returns the value of the nth item
func (l *LinkedList) ValueAtIndex(index int) *node2.Node {
	if l.IsEmpty() {
		return nil
	}
	start := 0
	for linked := l.node; linked != nil; linked = linked.Next {
		if index == start {
			return linked
		}
		start++
	}
	return nil
}

// Add an node at the cur
func (l *LinkedList) PushBack(node *node2.Node) {
	if l == nil {
		return
	}

	if endNode := l.Back(); endNode != nil {
		endNode.Next = node
	} else {
		l.node = node
	}
	l.num++
}

// Add an node at the Front
func (l *LinkedList) PushFront(node *node2.Node) {
	if l == nil {
		return
	}

	if l.node == nil {
		l.node = node
	} else {
		node.Next = l.node
		l.node = node
	}
	l.num++
}

// get value of cur node
func (l *LinkedList) Back() *node2.Node {
	if l.IsEmpty() {
		return nil
	}
	result := l.node
	for linked := l.node; linked != nil; linked = linked.Next {
		result = linked
	}
	return result
}

// get value of front node
func (l *LinkedList) Front() *node2.Node {
	if l.IsEmpty() {
		return nil
	}
	return l.node
}

func (l *LinkedList) PopFront() *node2.Node {
	if l.IsEmpty() {
		return nil
	}
	return l.Remove(0)
}

func (l *LinkedList) PopBack() *node2.Node {
	if l.IsEmpty() {
		return nil
	}

	return l.Remove(l.num - 1)
}

// bool returns true if is empty
func (l *LinkedList) IsEmpty() bool {
	if l == nil {
		return true
	}
	return l.num == 0
}

func (l *LinkedList) Remove(index int) *node2.Node {
	if l.IsEmpty() {
		return nil
	}

	var result *node2.Node
	if index == 0 {
		result := l.node
		l.node = result.Next
		l.num--
		return result
	} else {
		prevNode := l.ValueAtIndex(index - 1)
		result = l.removeNode(prevNode)
	}
	return result
}

func (l *LinkedList) RemoveFirstValue(node *node2.Node) {
	if l.IsEmpty() {
		return
	}
	if l.node.Equals(node) {
		l.Remove(0)
		return
	}
	var result *node2.Node
	for curNode := l.node; curNode != nil; curNode = curNode.Next {
		if curNode.Next.Equals(curNode) {
			result = curNode
			break
		}
	}
	l.removeNode(result)
}

func (l *LinkedList) removeNode(prevNode *node2.Node) *node2.Node {
	if prevNode == nil {
		return nil
	}
	removeNode := prevNode.Next
	prevNode.Next = removeNode.Next
	l.num--
	return removeNode
}
