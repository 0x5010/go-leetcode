package leetcode0894

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func allPossibleFBT(N int) []*TreeNode {
	if N == 1 {
		return []*TreeNode{&TreeNode{}}
	}
	res := []*TreeNode{}
	for i := 1; i < N-1; i += 2 {
		l, r := allPossibleFBT(i), allPossibleFBT(N-i-1)
		for _, ln := range l {
			for _, rn := range r {
				res = append(res, &TreeNode{
					Left:  ln,
					Right: rn,
				})
			}
		}
	}
	return res
}
