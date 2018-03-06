package leetcode0114

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
	var prev *TreeNode
	recur(root, &prev)
}

func recur(node *TreeNode, prev **TreeNode) {
	if node == nil {
		return
	}
	if *prev != nil {
		(*prev).Left = nil
		(*prev).Right = node
	}
	*prev = node
	l, r := node.Left, node.Right
	if l != nil {
		recur(l, prev)
	}
	if r != nil {
		recur(r, prev)
	}
}
