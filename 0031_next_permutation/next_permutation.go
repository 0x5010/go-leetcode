package leetcode0031

import (
	"sort"
)

func nextPermutation(nums []int) {
	l := len(nums)

	if l <= 1 {
		return
	}

	i := 0
	for i = l - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}

	if i > 0 {
		sort.Ints(nums[i:])
		for j := i - 1; j < l; j++ {
			if nums[j] > nums[i-1] {
				nums[i-1], nums[j] = nums[j], nums[i-1]
				return
			}
		}
	}
	sort.Ints(nums)
}
