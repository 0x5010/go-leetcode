package leetcode0863

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	res := []int{}
	var forChildren func(*TreeNode, int)
	forChildren = func(node *TreeNode, d int) {
		if node == nil {
			return
		}
		if d == K {
			res = append(res, node.Val)
			return
		}
		if d > K {
			return
		}
		d++
		forChildren(node.Left, d)
		forChildren(node.Right, d)
	}

	var distance func(*TreeNode) int
	distance = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		if node == target {
			forChildren(node, 0)
			return 0
		}
		l := distance(node.Left)
		if l != -1 {
			forChildren(node.Right, l+2)
			if l+1 == K {
				res = append(res, node.Val)
			}
			return l + 1
		}
		r := distance(node.Right)
		if r != -1 {
			forChildren(node.Left, r+2)
			if r+1 == K {
				res = append(res, node.Val)
			}
			return r + 1
		}
		return -1
	}
	distance(root)
	return res
}
