package leetcode0128

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	sort.Ints(nums)

	max, tmp := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			tmp++
		} else if nums[i] != nums[i-1] {
			if max < tmp {
				max = tmp
			}
			tmp = 1
		}
	}
	if max < tmp {
		return tmp
	}
	return max
}
