package leetcode0662

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func widthOfBinaryTree(root *TreeNode) int {
	l := [][2]int{}
	var dfs func(*TreeNode, int, int) int
	dfs = func(node *TreeNode, level, order int) int {
		if node == nil {
			return 0
		}
		if len(l) == level {
			l = append(l, [2]int{order, order})
		} else {
			l[level][1] = order
		}
		cur := l[level][1] - l[level][0] + 1
		l := dfs(node.Left, level+1, 2*order)
		r := dfs(node.Right, level+1, 2*order+1)
		return max(cur, max(l, r))
	}
	return dfs(root, 0, 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
