package leetcode0971

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	res := []int{}
	i := 0
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if node.Val != voyage[i] {
			i++
			return false
		}
		i++
		if node.Left != nil && node.Left.Val != voyage[i] {
			res = append(res, node.Val)
			return dfs(node.Right) && dfs(node.Left)
		}
		return dfs(node.Left) && dfs(node.Right)
	}
	if dfs(root) {
		return res
	}
	return []int{-1}
}
