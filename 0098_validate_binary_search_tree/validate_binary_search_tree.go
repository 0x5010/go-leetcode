package leetcode0098

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return validate(root, 0, 0, false, false)
}

func validate(node *TreeNode, low, high int, lowLimit, highLimit bool) bool {
	if node == nil {
		return true
	}
	if lowLimit && node.Val <= low {
		return false
	}
	if highLimit && node.Val >= high {
		return false
	}
	if !validate(node.Left, low, node.Val, lowLimit, true) {
		return false
	}
	return validate(node.Right, node.Val, high, true, highLimit)
}
