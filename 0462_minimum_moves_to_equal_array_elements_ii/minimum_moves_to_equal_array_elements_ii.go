package leetcode0462

import "sort"

func minMoves2(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	res := 0
	for i, j := 0, n-1; i < j; {
		res += nums[j] - nums[i]
		i++
		j--
	}
	return res
}
