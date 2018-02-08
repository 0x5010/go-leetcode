package leetcode0023

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l2.Next, l1)
	return l2

}

func mergeKLists(lists []*ListNode) *ListNode {
	llen := len(lists)
	if llen == 0 {
		return nil
	} else if llen == 1 {
		return lists[0]
	}
	half := llen / 2
	return mergeTwoLists(mergeKLists(lists[:half]), mergeKLists(lists[half:]))
}
