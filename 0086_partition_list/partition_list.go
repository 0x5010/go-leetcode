package leetcode0086

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n1, n2 := &ListNode{}, &ListNode{}
	p1, p2 := n1, n2
	p := head
	for p != nil {
		if p.Val < x {
			n1.Next = p
			n1 = n1.Next
		} else {
			n2.Next = p
			n2 = n2.Next
		}
		p, p.Next = p.Next, nil
	}
	n1.Next = p2.Next
	return p1.Next
}
