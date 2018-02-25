package leetcode0083

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	tmp := head
	for tmp.Next != nil {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}
