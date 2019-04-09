package leetcode0968

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minCameraCover(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 2
		}
		l, r := dfs(node.Left), dfs(node.Right)
		if l == 0 || r == 0 {
			res++
			return 1
		}
		if l == 1 || r == 1 {
			return 2
		}
		return 0
	}
	if dfs(root) < 1 {
		res++
	}
	return res
}
