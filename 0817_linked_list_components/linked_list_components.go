package leetcode0817

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func numComponents(head *ListNode, G []int) int {
	m := make(map[int]bool, len(G))
	for _, g := range G {
		m[g] = true
	}
	res := 0
	for head != nil {
		if m[head.Val] && (head.Next == nil || !m[head.Next.Val]) {
			res++
		}
		head = head.Next
	}
	return res
}
