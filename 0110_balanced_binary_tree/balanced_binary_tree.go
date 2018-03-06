package leetcode0110

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	_, res := check(root)
	return res
}

func check(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	ld, lb := check(root.Left)
	if !lb {
		return -1, false
	}
	rd, rb := check(root.Right)
	if !rb {
		return -1, false
	}

	if ld > 1+rd || rd > 1+ld {
		return -1, false
	}
	depth := 1 + ld
	if rd > ld {
		depth = 1 + rd
	}
	return depth, true
}
