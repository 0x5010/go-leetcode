package leetcode0002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func next(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
}

func add(n1, n2 *ListNode, i int) (sum, carry int) {
	if n1 != nil {
		sum += n1.Val
	}

	if n2 != nil {
		sum += n2.Val
	}

	sum += i
	if sum > 9 {
		carry = sum / 10
		sum %= 10
	}
	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	carry := 0

	for {
		var sum int
		sum, carry = add(l1, l2, carry)
		tmp.Val = sum

		l1, l2 = next(l1), next(l2)

		if l1 == nil && l2 == nil {
			break
		}

		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}

	if carry != 0 {
		tmp.Next = &ListNode{Val: carry}
	}
	return result
}
