package leetcode0016

import "sort"

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func threeSumClosest(nums []int, target int) int {
	diff := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 {
			diff = nums[0] + nums[1] + nums[len(nums)-1] - target
		} else if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			tmp := nums[i] + nums[l] + nums[r] - target
			if tmp == 0 {
				return target
			}

			if tmp < 0 {
				l++
			} else {
				r--
			}
			if abs(tmp) < abs(diff) {
				diff = tmp
			}
		}
	}
	return diff + target
}
