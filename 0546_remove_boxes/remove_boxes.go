package leetcode0546

func removeBoxes(boxes []int) int {
	n := len(boxes)
	if n < 2 {
		return n
	}
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
		for k := 0; k <= i; k++ {
			dp[i][i][k] = (k + 1) * (k + 1)
		}
	}

	for l := 1; l < n; l++ {
		for j := l; j < n; j++ {
			i := j - l
			for k := 0; k <= i; k++ {
				res := (k+1)*(k+1) + dp[i+1][j][0]
				for m := i + 1; m <= j; m++ {
					if boxes[m] == boxes[i] {
						res = max(res, dp[i+1][m-1][0]+dp[m][j][k+1])
					}
				}
				dp[i][j][k] = res
			}
		}
	}
	return dp[0][n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
