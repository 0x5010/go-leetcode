package leetcode0018

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				if l > j+1 && nums[l-1] == nums[l] {
					l++
					continue
				}
				tmp := nums[i] + nums[j] + nums[l] + nums[r]
				if tmp == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				} else if tmp < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
