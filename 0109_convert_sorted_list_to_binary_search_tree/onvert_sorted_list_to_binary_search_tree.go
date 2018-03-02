package leetcode00109

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	var prev *ListNode
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	p := slow
	if prev != nil {
		prev.Next = nil
	} else {
		head = nil
	}

	return &TreeNode{
		Val:   p.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(p.Next),
	}
}
