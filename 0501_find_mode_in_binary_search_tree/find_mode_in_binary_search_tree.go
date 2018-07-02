package leetcode0501

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findMode(root *TreeNode) []int {
	res := []int{}

	max, pre, cur := 0, 0, 0
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		if node.Val == pre {
			cur++
		} else {
			cur = 1
		}
		if cur == max {
			res = append(res, node.Val)
		} else if cur > max {
			max = cur
			res = []int{node.Val}
		}

		pre = node.Val
		inorder(node.Right)
	}

	inorder(root)
	return res
}
