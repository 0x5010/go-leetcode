package leetcode0611

import "sort"

func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	res := 0
	for i := n - 1; i >= 2; i-- {
		for j, k := 0, i-1; j < k; {
			if nums[j]+nums[k] > nums[i] {
				res += k - j
				k--
			} else {
				j++
			}
		}
	}
	return res
}
