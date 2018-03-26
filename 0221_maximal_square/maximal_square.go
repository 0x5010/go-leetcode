package leetcode0221

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([]int, m+1)
	maxEdge, pre := 0, 0
	for j := 0; j < n; j++ {
		for i := 1; i < m+1; i++ {
			tmp := dp[i]
			if matrix[i-1][j] == '1' {
				dp[i] = min(dp[i], min(dp[i-1], pre)) + 1
				if dp[i] > maxEdge {
					maxEdge = dp[i]
				}
			} else {
				dp[i] = 0
			}
			pre = tmp
		}
	}
	return maxEdge * maxEdge
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
