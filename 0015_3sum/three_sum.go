package leetcode0015

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			tmp := nums[i] + nums[l] + nums[r]
			if tmp == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if tmp < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
