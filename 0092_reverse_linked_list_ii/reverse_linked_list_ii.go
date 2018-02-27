package leetcode0092

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	tmp := &ListNode{-1, head}
	pre := tmp
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < n-m; i++ {
		move := cur.Next
		cur.Next = move.Next
		move.Next = pre.Next
		pre.Next = move
	}
	return tmp.Next
}
