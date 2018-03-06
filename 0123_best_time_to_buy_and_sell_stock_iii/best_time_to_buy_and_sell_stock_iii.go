package leetcode0123

import (
	"math"
)

func maxProfit(prices []int) int {
	b1, b2, s1, s2 := math.MinInt32, math.MinInt32, 0, 0
	for _, p := range prices {
		b1 = max(b1, -p)
		s1 = max(s1, b1+p)
		b2 = max(b2, s1-p)
		s2 = max(s2, b2+p)
	}
	return s2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
