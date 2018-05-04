package leetcode0382

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "math/rand"

type Solution struct {
	head *ListNode
	n    int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	count := 0
	tmp := head
	for tmp != nil {
		count++
		tmp = tmp.Next
	}
	return Solution{
		head: head,
		n:    count,
	}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	tmp := this.head
	for n := rand.Intn(this.n); n > 0; n-- {
		tmp = tmp.Next
	}
	return tmp.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
