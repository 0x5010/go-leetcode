package leetcode0337

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		lRob, lNotRob := dfs(node.Left)
		rRob, rNotRob := dfs(node.Right)
		return node.Val + lNotRob + rNotRob, max(lRob, lNotRob) + max(rRob, rNotRob)
	}
	return max(dfs(root))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
