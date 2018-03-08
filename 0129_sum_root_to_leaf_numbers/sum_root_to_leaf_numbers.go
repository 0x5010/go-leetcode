package leetcode0129

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumNumbers(root *TreeNode) int {
	res := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, tmp int) {
		if node == nil {
			return
		}

		tmp = tmp*10 + node.Val
		if node.Left == nil && node.Right == nil {
			res += tmp
			return
		}
		dfs(node.Left, tmp)
		dfs(node.Right, tmp)
	}
	dfs(root, 0)
	return res
}
