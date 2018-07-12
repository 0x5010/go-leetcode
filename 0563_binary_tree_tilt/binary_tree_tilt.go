package leetcode0563

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findTilt(root *TreeNode) int {
	res := 0
	var postorder func(*TreeNode) int
	postorder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := postorder(node.Left), postorder(node.Right)
		res += abs(l, r)
		return l + node.Val + r
	}
	postorder(root)
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
