package leetcode0375

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for j := 2; j < n+1; j++ {
		for i := j - 1; i > 0; i-- {
			dp[i][j] = i + dp[i+1][j]
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], k+max(dp[i][k-1], dp[k+1][j]))
			}
		}
	}
	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
