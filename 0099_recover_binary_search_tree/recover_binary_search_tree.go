package leetcode0099

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recoverTree(root *TreeNode) {
	var n1, n2, prev *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			dfs(node.Left)
		}

		if prev != nil && prev.Val > node.Val {
			if n1 == nil {
				n1 = prev
			}
			if n1 != nil {
				n2 = node
			}
		}
		prev = node

		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)
	n1.Val, n2.Val = n2.Val, n1.Val
}
