package leetcode0652

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	count := map[string]int{}
	res := []*TreeNode{}
	var dfs func(*TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "*"
		}
		l, r := dfs(node.Left), dfs(node.Right)
		k := strconv.Itoa(node.Val) + l + r
		count[k]++
		if count[k] == 2 {
			res = append(res, node)
		}
		return k
	}
	dfs(root)
	return res
}
