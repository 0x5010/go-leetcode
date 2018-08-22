package leetcode0818

import "math"

func racecar(target int) int {
	dp := make([]int, target+1)

	var recur func(int) int
	recur = func(t int) int {
		if dp[t] > 0 {
			return dp[t]
		}
		n := uint(math.Log2(float64(t))) + 1
		if t == 1<<n-1 {
			dp[t] = int(n)
		} else {
			dp[t] = recur(1<<n-1-t) + int(n) + 1
			for m := uint(0); m < n-1; m++ {
				dp[t] = min(dp[t], recur(t-1<<(n-1)+1<<m)+int(n-1)+1+int(m)+1)
			}
		}
		return dp[t]
	}
	return recur(target)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
