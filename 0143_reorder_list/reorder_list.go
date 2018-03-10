package leetcode0143

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	p1, p2 := head, head.Next
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	cur := p1.Next
	p1.Next = nil

	p2 = cur.Next
	cur.Next = nil

	for p2 != nil {
		p1 = p2.Next
		p2.Next = cur
		cur = p2
		p2 = p1
	}
	p1 = head
	p2 = cur
	for p1 != nil {
		tmp := p1.Next
		p1.Next = p2
		p1 = p2
		p2 = tmp
	}
}
