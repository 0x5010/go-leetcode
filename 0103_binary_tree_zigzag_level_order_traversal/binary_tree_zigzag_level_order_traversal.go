package leetcode00103

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	s := []*TreeNode{root}
	res := [][]int{}
	zigzag := false
	for len(s) != 0 {
		tmp := []*TreeNode{}
		l := []int{}
		for i, node := range s {
			if zigzag {
				l = append(l, s[len(s)-i-1].Val)
			} else {
				l = append(l, node.Val)
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		res = append(res, l)
		s = tmp
		zigzag = !zigzag
	}
	return res
}
