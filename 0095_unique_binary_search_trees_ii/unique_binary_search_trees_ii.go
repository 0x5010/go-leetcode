package leetcode0095

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return nil
	}
	return genTrees(1, n)
}

func genTrees(start, end int) []*TreeNode {
	list := []*TreeNode{}

	if start > end {
		return append(list, nil)
	} else if start == end {
		return append(list, &TreeNode{start, nil, nil})
	}

	for i := start; i <= end; i++ {
		l := genTrees(start, i-1)
		r := genTrees(i+1, end)

		for _, lnode := range l {
			for _, rnode := range r {
				list = append(list, &TreeNode{i, lnode, rnode})
			}
		}
	}
	return list
}
