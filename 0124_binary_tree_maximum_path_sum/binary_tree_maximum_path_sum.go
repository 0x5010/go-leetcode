package leetcode0124

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	max := root.Val

	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l, r := dfs(root.Left), dfs(root.Right)
		if l < 0 {
			l = 0
		}
		if r < 0 {
			r = 0
		}
		sum := l + root.Val + r
		if sum > max {
			max = sum
		}
		if l > r {
			return l + root.Val
		}
		return root.Val + r
	}
	dfs(root)
	return max
}
