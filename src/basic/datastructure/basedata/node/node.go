package node

import "reflect"

type Node struct {
	Value interface{}
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value}
}

func (e *Node) Equals(node *Node) bool {
	if e == nil || node == nil {
		return false
	}
	return reflect.DeepEqual(e.Value, node.Value)
}

