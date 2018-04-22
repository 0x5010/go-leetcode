package leetcode0334

import "math"

func increasingTriplet(nums []int) bool {
	i, j := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= i {
			i = num
		} else if num <= j {
			j = num
		} else {
			return true
		}
	}
	return false
}
