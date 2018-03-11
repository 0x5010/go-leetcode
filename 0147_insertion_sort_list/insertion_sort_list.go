package leetcode0147

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertionSortList(head *ListNode) *ListNode {
	h := &ListNode{Next: head}

	cur := head
	for cur != nil && cur.Next != nil {
		p := cur.Next
		if cur.Val <= p.Val {
			cur = p
			continue
		}

		cur.Next = p.Next
		pre, next := h, h.Next
		for next.Val < p.Val {
			pre = next
			next = next.Next
		}
		pre.Next = p
		p.Next = next
	}
	return h.Next
}
