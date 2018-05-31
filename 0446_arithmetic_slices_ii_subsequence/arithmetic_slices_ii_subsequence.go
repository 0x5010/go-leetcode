package leetcode0446

import (
	"math"
)

func numberOfArithmeticSlices(A []int) int {
	res := 0
	n := len(A)
	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := A[i] - A[j]
			if diff <= math.MinInt32 || diff > math.MaxInt32 {
				continue
			}
			dp[i][diff] += dp[j][diff] + 1
			res += dp[j][diff]
		}
	}
	return res
}
