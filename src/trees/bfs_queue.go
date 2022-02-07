package trees

import (
	"fmt"
	"github.com/play-code/src/basedata/ele"
	"github.com/play-code/src/basedata/tree"
	"github.com/play-code/src/container/queue"
)

func BfsQueue(root *tree.TreeNode) {
	q := queue.NewQueue()
	q.Enqueue(ele.NewEle(root))
	if !q.IsEmpty() {
		treeNode := q.Dequeue().Value().(*tree.TreeNode)
		fmt.Println(treeNode.Val)
		if treeNode.Left != nil {
			q.Enqueue(ele.NewEle(treeNode.Left))
		}
		if treeNode.Right != nil {
			q.Enqueue(ele.NewEle(treeNode.Right))
		}
	}

}
