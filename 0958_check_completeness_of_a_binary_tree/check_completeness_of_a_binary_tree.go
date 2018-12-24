package leetcode0958

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	i := 0
	for i < len(queue) && queue[i] != nil {
		node := queue[i]
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
		i++
	}
	for i < len(queue) && queue[i] == nil {
		i++
	}
	return i == len(queue)
}
