package leetcode0872

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	s1, s2 := []*TreeNode{root1}, []*TreeNode{root2}
	for len(s1) != 0 && len(s2) != 0 {
		if dfs(&s1) != dfs(&s2) {
			return false
		}
	}
	return len(s1) == 0 && len(s2) == 0
}

func dfs(s *[]*TreeNode) int {
	for {
		node := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		if node.Left != nil {
			*s = append(*s, node.Left)
		}
		if node.Right != nil {
			*s = append(*s, node.Right)
		}
		if node.Left == nil && node.Right == nil {
			return node.Val
		}
	}
}
