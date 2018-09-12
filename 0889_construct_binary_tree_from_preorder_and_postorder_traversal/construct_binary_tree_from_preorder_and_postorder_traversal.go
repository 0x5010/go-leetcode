package leetcode0889

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func constructFromPrePost(pre []int, post []int) *TreeNode {
	n := len(pre)
	if n == 0 {
		return nil
	}
	res := &TreeNode{
		Val: pre[0],
	}
	if n == 1 {
		return res
	}
	l := pre[1]
	i := 0
	for i < n {
		if post[i] == l {
			break
		}
		i++
	}
	res.Left = constructFromPrePost(pre[1:i+2], post[:i+1])
	res.Right = constructFromPrePost(pre[i+2:], post[i+1:n-1])
	return res
}
