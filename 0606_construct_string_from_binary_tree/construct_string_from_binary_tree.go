package leetcode0606

import (
	"bytes"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func tree2str(t *TreeNode) string {
	bs := bytes.Buffer{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		bs.WriteString(strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil {
			return
		}
		bs.WriteString("(")
		if node.Left != nil {
			dfs(node.Left)
		}
		bs.WriteString(")")
		if node.Right != nil {
			bs.WriteString("(")
			dfs(node.Right)
			bs.WriteString(")")
		}
	}
	dfs(t)
	return bs.String()
}
