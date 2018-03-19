package leetcode0188

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	if k >= n/2 {
		res := 0
		for i := 1; i < n; i++ {
			tmp := prices[i] - prices[i-1]
			if tmp > 0 {
				res += tmp
			}
		}
		return res
	}

	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, n)
	}
	for i := 1; i < k+1; i++ {
		tmp := -prices[0]
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+tmp)
			tmp = max(tmp, dp[i-1][j-1]-prices[j])
		}
	}
	return dp[k][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
