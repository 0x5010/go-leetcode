package leetcode0234

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	var rev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		rev = &ListNode{Val: slow.Val, Next: rev}
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	for rev != nil && rev.Val == slow.Val {
		slow = slow.Next
		rev = rev.Next
	}
	return rev == nil
}
