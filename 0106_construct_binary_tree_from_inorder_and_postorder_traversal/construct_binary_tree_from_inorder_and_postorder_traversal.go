package leetcode0106

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	res := &TreeNode{Val: postorder[len(postorder)-1]}
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

	res.Left = buildTree(inorder[:index], postorder[:index])
	res.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return res
}
