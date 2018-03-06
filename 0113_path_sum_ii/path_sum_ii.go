package leetcode0113

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	var recur func(*TreeNode, int, []int)
	recur = func(node *TreeNode, sum int, path []int) {
		if node == nil {
			return
		}
		sum -= node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			if sum == 0 {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
				return
			}
		}
		recur(node.Left, sum, path)
		recur(node.Right, sum, path)

	}
	recur(root, sum, []int{})
	return res
}
