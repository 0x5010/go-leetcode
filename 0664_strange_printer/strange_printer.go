package leetcode0664

func strangePrinter(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for k := 2; k <= n; k++ {
		for i := 0; i+k-1 < n; i++ {
			j := i + k - 1
			dp[i][j] = dp[i][j-1] + 1
			if s[j-1] == s[j] {
				dp[i][j]--
			}

			for k := i + 1; k <= j; k++ {
				if s[k-1] == s[j] {
					dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k][j]-1)
				}
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
