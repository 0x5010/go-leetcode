package leetcode0148

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	fast = slow.Next
	slow.Next = nil

	l1 := sortList(head)
	l2 := sortList(fast)

	h := &ListNode{}
	rear := h

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			rear.Next = l1
			l1 = l1.Next
		} else {
			rear.Next = l2
			l2 = l2.Next
		}
		rear = rear.Next
		rear.Next = nil
	}
	if l1 != nil {
		rear.Next = l1
	} else if l2 != nil {
		rear.Next = l2
	}
	return h.Next
}
