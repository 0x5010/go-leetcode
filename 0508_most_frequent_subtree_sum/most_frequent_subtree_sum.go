package leetcode0508

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findFrequentTreeSum(root *TreeNode) []int {
	m := map[int]int{}
	max := 0

	var postorder func(*TreeNode) int
	postorder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + postorder(node.Left) + postorder(node.Right)

		m[sum]++
		if m[sum] > max {
			max = m[sum]
		}
		return sum
	}

	postorder(root)

	res := []int{}
	for sum, count := range m {
		if count == max {
			res = append(res, sum)
		}
	}
	return res
}
