package leetcode0965

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isUnivalTree(root *TreeNode) bool {
	return (root.Left == nil || root.Left.Val == root.Val && isUnivalTree(root.Left)) &&
		(root.Right == nil || root.Right.Val == root.Val && isUnivalTree(root.Right))
}
