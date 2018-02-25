package leetcode0082

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		val := head.Val
		head = head.Next.Next

		for head != nil && head.Val == val {
			head = head.Next
		}
		return deleteDuplicates(head)
	}
	head.Next = deleteDuplicates(head.Next)
	return head
}
