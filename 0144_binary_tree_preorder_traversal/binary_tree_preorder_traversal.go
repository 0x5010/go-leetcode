package leetcode0144

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []*TreeNode{}
	p := root
	for len(stack) != 0 || p != nil {
		for p != nil {
			res = append(res, p.Val)
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	return res
}
