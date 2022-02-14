package trees

import (
	"github.com/play-code/src/basic/datastructure/basedata/ele"
	"github.com/play-code/src/basic/datastructure/basedata/tree"
	"github.com/play-code/src/basic/datastructure/container/queue"
)

func BfsQueue(root *tree.TreeNode) []ele.TreeEle {
	result := make([]ele.TreeEle, 0, 10)
	q := queue.NewQueue()
	q.Enqueue(ele.NewEle(root))
	for !q.IsEmpty() {
		treeNode := q.Dequeue().Value().(*tree.TreeNode)
		result = append(result, treeNode.Val)
		if treeNode.Left != nil {
			q.Enqueue(ele.NewEle(treeNode.Left))
		}
		if treeNode.Right != nil {
			q.Enqueue(ele.NewEle(treeNode.Right))
		}
	}

	return result
}
