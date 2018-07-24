package leetcode0655

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func printTree(root *TreeNode) [][]string {
	depth := getDepth(root, 0) + 1
	width := 1<<uint(depth) - 1
	res := make([][]string, depth)
	for i := range res {
		res[i] = make([]string, width)
	}
	var recur func(*TreeNode, int, int)
	recur = func(node *TreeNode, level, bin int) {
		if node == nil {
			return
		}
		span := width >> uint(level)
		res[level][span*bin+bin+span/2] = strconv.Itoa(node.Val)
		recur(node.Left, level+1, 2*bin)
		recur(node.Right, level+1, 2*bin+1)
	}
	recur(root, 0, 0)
	return res
}

func getDepth(node *TreeNode, d int) int {
	l := d
	if node.Left != nil {
		l = getDepth(node.Left, d+1)
	}
	r := d
	if node.Right != nil {
		r = getDepth(node.Right, d+1)
	}
	if l >= r {
		return l
	}
	return r
}
