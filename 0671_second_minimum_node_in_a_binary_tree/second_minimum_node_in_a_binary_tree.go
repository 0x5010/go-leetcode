package leetcode0671

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findSecondMinimumValue(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, first int) int {
		if node == nil {
			return -1
		}
		if node.Val != first {
			return node.Val
		}
		l, r := dfs(node.Left, first), dfs(node.Right, first)
		if l == -1 {
			return r
		}
		if r == -1 {
			return l
		}
		return min(l, r)
	}
	return dfs(root, root.Val)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
