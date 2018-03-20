package leetcode0199

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	view(root, &res, 0)
	return res
}

func view(node *TreeNode, res *[]int, depth int) {
	if node == nil {
		return
	}
	if depth == len(*res) {
		*res = append(*res, node.Val)
	}
	depth++
	view(node.Right, res, depth)
	view(node.Left, res, depth)
}
