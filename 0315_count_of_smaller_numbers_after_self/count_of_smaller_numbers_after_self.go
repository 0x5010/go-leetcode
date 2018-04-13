package leetcode0315

func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	var root *node
	for i := len(nums) - 1; i >= 0; i-- {
		root = insert(root, nums[i], i, 0, &res)
	}
	return res
}

type node struct {
	l, r          *node
	val, sum, dup int
}

func insert(root *node, num, i, preSum int, res *[]int) *node {
	if root == nil {
		(*res)[i] = preSum
		return &node{
			val: num,
			dup: 1,
		}
	}

	if num == root.val {
		root.dup++
		(*res)[i] = preSum + root.sum
	} else if num < root.val {
		root.sum++
		root.l = insert(root.l, num, i, preSum, res)
	} else {
		root.r = insert(root.r, num, i, preSum+root.dup+root.sum, res)
	}
	return root
}
