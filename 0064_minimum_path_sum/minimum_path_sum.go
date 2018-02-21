package leetcode0064

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			dp[0][0] = grid[0][0]
		} else {
			dp[i][0] = grid[i][0] + dp[i-1][0]
		}
	}

	for i := 1; i < n; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			min := dp[i-1][j]
			if dp[i][j-1] < min {
				min = dp[i][j-1]
			}
			dp[i][j] = grid[i][j] + min
		}
	}

	return dp[m-1][n-1]
}
