package leetcode0623

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}
	if d == 2 {
		root.Left = &TreeNode{Val: v, Left: root.Left}
		root.Right = &TreeNode{Val: v, Right: root.Right}

	} else {
		if root.Left != nil {
			root.Left = addOneRow(root.Left, v, d-1)
		}
		if root.Right != nil {
			root.Right = addOneRow(root.Right, v, d-1)
		}
	}
	return root
}
