package leetcode0865

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var deep func(*TreeNode) (*TreeNode, int)
	deep = func(node *TreeNode) (*TreeNode, int) {
		if node == nil {
			return nil, 0
		}
		l, ld := deep(node.Left)
		r, rd := deep(node.Right)

		if ld == rd {
			return node, ld + 1
		} else if ld > rd {
			return l, ld + 1
		} else {
			return r, rd + 1
		}
	}
	res, _ := deep(root)
	return res
}
