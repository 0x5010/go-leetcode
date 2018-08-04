package leetcode0725

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func splitListToParts(root *ListNode, k int) []*ListNode {
	n := 0
	for node := root; node != nil; node = node.Next {
		n++
	}
	l, r := n/k, n%k
	res := make([]*ListNode, k)
	var prev *ListNode
	node := root
	for i := 0; node != nil && i < k; i, r = i+1, r-1 {
		res[i] = node
		ll := l
		if r > 0 {
			ll++
		}
		for j := 0; j < ll; j++ {
			prev = node
			node = node.Next
		}
		prev.Next = nil
	}
	return res
}
