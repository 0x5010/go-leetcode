package leetcode0897

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func increasingBST(root *TreeNode) *TreeNode {
	var inorder func(*TreeNode, *TreeNode) *TreeNode
	inorder = func(node, tail *TreeNode) *TreeNode {
		if node == nil {
			return tail
		}
		res := inorder(node.Left, node)
		node.Left = nil
		node.Right = inorder(node.Right, tail)
		return res
	}
	return inorder(root, nil)
}
