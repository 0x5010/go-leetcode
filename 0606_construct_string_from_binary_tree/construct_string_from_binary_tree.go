package leetcode0606

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	"strconv"
	"strings"
)

func tree2str(t *TreeNode) string {
	bs := strings.Builder{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		bs.WriteString(strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil {
			return
		}
		bs.WriteByte('(')
		if node.Left != nil {
			dfs(node.Left)
		}
		bs.WriteByte(')')
		if node.Right != nil {
			bs.WriteByte('(')
			dfs(node.Right)
			bs.WriteByte(')')
		}
	}
	dfs(t)
	return bs.String()
}
