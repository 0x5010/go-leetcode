package leetcode0230

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	count := 0
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil || count == k {
			return
		}
		preorder(node.Left)
		count++
		if count == k {
			res = node.Val
		}
		preorder(node.Right)
	}
	preorder(root)
	return res
}
