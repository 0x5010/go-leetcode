package leetcode0530

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt32
	var prev *int

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		if prev != nil {
			res = min(res, node.Val-*prev)
		}
		prev = &node.Val
		inorder(node.Right)
	}
	inorder(root)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
