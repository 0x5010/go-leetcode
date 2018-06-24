package leetcode0486

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i, num := range nums {
		dp[i] = make([]int, n)
		dp[i][i] = num
	}

	for k := 2; k <= n; k++ {
		for i := 0; i <= n-k; i++ {
			j := i + k - 1
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
