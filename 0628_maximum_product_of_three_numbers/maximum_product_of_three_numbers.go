package leetcode0628

import "math"

func maximumProduct(nums []int) int {
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32

	for _, n := range nums {
		if n > max1 {
			max3, max2, max1 = max2, max1, n
		} else if n > max2 {
			max3, max2 = max2, n
		} else if n > max3 {
			max3 = n
		}

		if n < min1 {
			min2, min1 = min1, n
		} else if n < min2 {
			min2 = n
		}
	}
	return max(max2*max3, min1*min2) * max1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
