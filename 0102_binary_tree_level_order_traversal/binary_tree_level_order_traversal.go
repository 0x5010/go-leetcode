package leetcode0102

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	s := []*TreeNode{root}
	res := [][]int{}
	for len(s) != 0 {
		tmp := []*TreeNode{}
		l := []int{}
		for _, node := range s {
			l = append(l, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		res = append(res, l)
		s = tmp
	}
	return res
}
