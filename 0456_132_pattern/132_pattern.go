package leetcode456

import "math"

func find132pattern(nums []int) bool {
	n := len(nums)
	s := math.MinInt32
	stack := []int{}
	for i := n - 1; i >= 0; i-- {
		if nums[i] < s {
			return true
		}
		for len(stack) != 0 && nums[i] > stack[len(stack)-1] {
			s = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}
