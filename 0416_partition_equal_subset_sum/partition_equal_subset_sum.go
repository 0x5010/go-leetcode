package leetcode0416

import (
	"sort"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	n := len(nums)
	var can func(int, int) bool
	can = func(index, currSum int) bool {
		if n <= index {
			return false
		}
		tmpSum := currSum + nums[index]
		if tmpSum == sum {
			return true
		}
		if tmpSum > sum {
			return false
		}
		index++
		return can(index, tmpSum) || can(index, currSum)
	}
	return can(0, 0)
}
