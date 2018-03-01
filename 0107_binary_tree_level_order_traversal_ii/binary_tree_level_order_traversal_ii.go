package leetcode00107

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrderBottom(root *TreeNode) [][]int {
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
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}
