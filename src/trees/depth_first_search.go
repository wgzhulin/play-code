package trees

import (
	"github.com/play-code/src/basedata/ele"
	"github.com/play-code/src/basedata/node"
	"github.com/play-code/src/basedata/tree"
	"github.com/play-code/src/container/list"
)

// 后进先出
func Dfs(root *tree.TreeNode) []ele.TreeEle {
	result := make([]ele.TreeEle, 0, 10)
	l := list.NewLinkedList()
	l.PushFront(node.NewNode(root))
	for !l.IsEmpty() {
		treeNode := l.PopBack().Value.(*tree.TreeNode)
		result = append(result, treeNode.Val)
		if treeNode.Right != nil {
			l.PushBack(node.NewNode(treeNode.Right))
		}
		if treeNode.Left != nil {
			l.PushBack(node.NewNode(treeNode.Left))
		}
	}

	return result
}
