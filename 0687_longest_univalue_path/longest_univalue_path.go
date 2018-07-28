package leetcode0687

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func longestUnivaluePath(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		lp, rp := 0, 0
		if node.Left != nil && node.Val == node.Left.Val {
			lp = l + 1
		}

		if node.Right != nil && node.Val == node.Right.Val {
			rp = r + 1
		}
		res = max(res, lp+rp)
		if lp > rp {
			return lp
		}
		return rp
	}
	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
