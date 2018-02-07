package leetcode0019

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getNth(head *ListNode, n int) (pnode *ListNode, headIsNthFromEnd bool) {
	pnode = head
	for head != nil {
		if n < 0 {
			pnode = pnode.Next
		}
		n--
		head = head.Next
	}
	headIsNthFromEnd = n == 0
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pnode, headIsNthFromEnd := getNth(head, n)
	if headIsNthFromEnd {
		return head.Next
	}
	pnode.Next = pnode.Next.Next
	return head
}
