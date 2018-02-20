package leetcode0061

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	first := head
	for i := 0; i < k; i++ {
		if first.Next == nil {
			return rotateRight(head, k%(i+1))
		}
		first = first.Next
	}

	second := head
	for first.Next != nil {
		second, first = second.Next, first.Next
	}

	newHead := second.Next
	second.Next, first.Next = nil, head

	return newHead
}
