package leetcode00105

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	res := &TreeNode{Val: preorder[0]}
	if len(inorder) == 1 {
		return res
	}

	index := 0
	for i, v := range inorder {
		if v == res.Val {
			index = i
			break
		}
	}

	res.Left = buildTree(preorder[1:index+1], inorder[:index])
	res.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return res
}
