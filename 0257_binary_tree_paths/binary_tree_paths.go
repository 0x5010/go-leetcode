package leetcode0257

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	res := []string{}
	var dfs func(string, *TreeNode)
	dfs = func(pre string, node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			res = append(res, pre+strconv.Itoa(node.Val))
			return
		}

		pre += strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			dfs(pre, node.Left)
		}
		if node.Right != nil {
			dfs(pre, node.Right)
		}

	}
	dfs("", root)
	return res
}
