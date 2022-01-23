package container

import (
	"reflect"
)

// Header
type LinkedList struct {
	num int
	node *Node
}

type Node struct {
	Value interface{}
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value}
}

func (e *Node) equal(node *Node) bool {
	if e == nil || node == nil {
		return false
	}
	return reflect.DeepEqual(e.Value, node.Value)
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) size() int {
	if l.isEmpty() {
		return 0
	}
	return l.num
}

func (l *LinkedList) insert(index int, node *Node) {
	if l.isEmpty() || index > l.num-1 {
		return
	}
	if index == 0 {
		l.pushFront(node)
	} else {
		indexnode := l.valueAtIndex(index - 1)
		node.Next = indexnode.Next.Next
		indexnode.Next = node
	}
	l.num++
}

// returns the value of the nth item
func (l *LinkedList) valueAtIndex(index int) *Node {
	if l.isEmpty() {
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

// add an node at the cur
func (l *LinkedList) pushBack(node *Node) {
	if l == nil {
		return
	}

	if lastnode := l.back(); lastnode != nil {
		lastnode.Next = node
	} else {
		l.node = node
	}
	l.num++
}

// add an node at the front
func (l *LinkedList) pushFront(node *Node) {
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
func (l *LinkedList) back() *Node {
	if l.isEmpty() {
		return nil
	}
	result := l.node
	for linked := l.node; linked != nil; linked = linked.Next {
		result = linked
	}
	return result
}

// get value of front node
func (l *LinkedList) front() *Node {
	if l.isEmpty() {
		return nil
	}
	return l.node
}

func (l *LinkedList) popFront() *Node {
	if l.isEmpty() {
		return nil
	}
	return l.remove(0)
}

func (l *LinkedList) popBack() *Node {
	if l.isEmpty() {
		return nil
	}
	return l.remove(l.num - 1)
}

// bool returns true if isEmpty
func (l *LinkedList) isEmpty() bool {
	if l == nil {
		return true
	}
	return l.num == 0
}

func (l *LinkedList) remove(index int) *Node {
	if l.isEmpty() {
		return nil
	}

	var result *Node
	if index == 0 {
		result := l.node
		l.node = result.Next
		l.num--
		return result
	} else {
		prevNode := l.valueAtIndex(index - 1)
		result = l.removeNode(prevNode)
	}
	return result
}

func (l *LinkedList) removeFirstValue(node *Node) {
	if l.isEmpty() {
		return
	}
	if l.node.equal(node) {
		l.remove(0)
		return
	}
	var prevNode *Node
	for node := l.node; node != nil; node = node.Next {
		if node.Next.equal(node) {
			prevNode = node
			break
		}
	}
	l.removeNode(prevNode)
}

func (l *LinkedList) removeNode(prevNode *Node) *Node {
	if prevNode == nil {
		return nil
	}
	removeNode := prevNode.Next
	prevNode.Next = removeNode.Next
	l.num--
	return removeNode
}
