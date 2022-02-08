package testdata

import (
	"github.com/play-code/src/basedata/ele"
	"github.com/play-code/src/basedata/tree"
)

// NewBinaryTreeData([]ele.TreeEle{2, 1, 3, 9, 0, 5, 4, 8})
func TreeNodeData() *tree.TreeNode {
	root := tree.NewTreeNode(2)

	treeNode3 := tree.NewTreeNode(3)
	treeNode3.Left = tree.NewTreeNode(5)
	treeNode3.Right = tree.NewTreeNode(4)

	treeNode9 := tree.NewTreeNode(9)
	treeNode9.Left = tree.NewTreeNode(8)

	treeNode1 := tree.NewTreeNode(1)
	treeNode1.Right = treeNode9

	root.Left = treeNode1
	root.Right = treeNode3

	return root
}

// 切片数据转为二叉树结构
func NewBinaryTreeData(input []ele.TreeEle) *tree.TreeNode {
	s := make([]ele.TreeEle, len(input)+1)
	copy(s[1:], input)
	return binaryTree(s, 1)
}

func binaryTree(input []ele.TreeEle, index int) *tree.TreeNode {
	if index < len(input) {
		val := input[index]
		if val != 0 {
			return &tree.TreeNode{
				Val:   input[index],
				Left:  binaryTree(input, index*2),
				Right: binaryTree(input, index*2+1),
			}
		}
	}
	return nil
}
