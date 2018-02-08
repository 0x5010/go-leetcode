package leetcode0025

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKNodes(head *ListNode, k int) *ListNode {
	end := head
	for k > 0 && end != nil {
		end = end.Next
		k--
	}
	if k > 0 {
		return head
	}
	res, pNode := end, head
	var qNode *ListNode
	for pNode != end {
		qNode = pNode.Next
		pNode.Next = res
		res = pNode
		pNode = qNode
	}
	return res
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 0 {
		return head
	}
	tmp := &ListNode{Next: head}
	p := tmp
	for p != nil {
		p.Next = reverseKNodes(p.Next, k)
		for i := 0; p != nil && i < k; i++ {
			p = p.Next
		}
	}
	return tmp.Next
}
