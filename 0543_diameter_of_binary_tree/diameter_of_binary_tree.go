package leetcode0543

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var maxDepth func(*TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := maxDepth(node.Left), maxDepth(node.Right)
		res = max(res, l+r)
		return max(l, r) + 1
	}
	maxDepth(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
