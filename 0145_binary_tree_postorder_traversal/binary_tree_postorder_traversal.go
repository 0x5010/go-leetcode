package leetcode0145

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []*TreeNode{root}
	p := root
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		l, r := node.Left, node.Right
		if l == p || r == p || l == nil && r == nil {
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
			p = node
		} else {
			if r != nil {
				stack = append(stack, r)
			}
			if l != nil {
				stack = append(stack, l)
			}
		}
	}
	return res
}
