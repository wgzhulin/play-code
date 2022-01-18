package code

/*
144. 二叉树的前序遍历
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

示例 1：
输入：root = [1,null,2,3]
输出：[1,2,3]

https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
*/

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var tmp []int
	tmp = append(tmp, root.Val)
	if root.Left != nil {
		tmp = append(tmp, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		tmp = append(tmp, preorderTraversal(root.Right)...)
	}
	return tmp
}