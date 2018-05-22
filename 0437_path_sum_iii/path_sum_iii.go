package leetcode0437

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, sum int) int {
	m := make(map[int]int)
	m[0] = 1

	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, curSum int) int {
		if node == nil {
			return 0
		}
		curSum += node.Val
		res := m[curSum-sum]
		m[curSum]++
		res += dfs(node.Left, curSum) + dfs(node.Right, curSum)
		m[curSum]--
		return res
	}
	return dfs(root, 0)
}
