package leetcode0654

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func constructMaximumBinaryTree(nums []int) *TreeNode {
	s := []*TreeNode{}
	for _, num := range nums {
		node := &TreeNode{Val: num}
		for len(s) != 0 && s[len(s)-1].Val < num {
			node.Left = s[len(s)-1]
			s = s[:len(s)-1]
		}
		if len(s) != 0 {
			s[len(s)-1].Right = node
		}
		s = append(s, node)
	}
	return s[0]
}
