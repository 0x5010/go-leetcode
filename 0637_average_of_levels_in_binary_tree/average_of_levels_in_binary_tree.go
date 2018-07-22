package leetcode0637

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		n := len(nodes)
		sum := 0
		for _, node := range nodes {
			sum += node.Val
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		res = append(res, float64(sum)/float64(n))
		nodes = nodes[n:]
	}
	return res
}
