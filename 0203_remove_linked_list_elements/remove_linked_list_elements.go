package leetcode0203

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	h := &ListNode{Next: head}

	tmp := h
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return h.Next
}
