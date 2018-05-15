package leetcode0414

import "math"

func thirdMax(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num == max1 || num == max2 {
			continue
		}

		if num > max1 {
			max1, max2, max3 = num, max1, max2
		} else if num > max2 {
			max2, max3 = num, max2
		} else if num > max3 {
			max3 = num
		}
	}
	if max3 == math.MinInt64 {
		return max1
	}
	return max3
}
