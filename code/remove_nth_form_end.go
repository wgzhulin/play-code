package code

/*
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。



示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sum := 0
	for node := head; node != nil; node = node.Next {
		sum++
	}

	if sum-n-1 < 0 {
		head = head.Next
		return head
	}

	cur := 0
	for node := head; node != nil; node = node.Next {
		if cur == sum-n-1 {
			node.Next = node.Next.Next
			break
		}
		cur++
	}

	return head
}
