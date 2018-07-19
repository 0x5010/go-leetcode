package leetcode0629

func kInversePairs(n int, k int) int {
	m := 1000000007
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		maxJ := min(k, i*(i-1)/2)
		for j := 1; j <= maxJ; j++ {
			dp[i][j] = (dp[i][j-1] + dp[i-1][j]) % m
			if j >= i {
				dp[i][j] -= dp[i-1][j-i]
				if dp[i][j] < 0 {
					dp[i][j] += m
				}
			}
		}
	}

	return dp[n][k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
