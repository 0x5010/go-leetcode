package leetcode0741

func cherryPickup(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[0][0] = grid[0][0]

	pathLen := n + n - 1
	for pl := 2; pl <= pathLen; pl++ {
		for x1 := n - 1; x1 >= 0; x1-- {
			for x2 := x1; x2 >= 0; x2-- {
				y1 := pl - 1 - x1
				y2 := pl - 1 - x2
				if y1 < 0 || y2 < 0 || y1 >= n || y2 >= n {
					continue
				}
				if grid[y1][x1] < 0 || grid[y2][x2] < 0 {
					dp[x1][x2] = -1
					continue
				}

				best := -1
				delta := grid[y1][x1]

				if x1 != x2 {
					delta += grid[y2][x2]
				}
				if x1 > 0 && x2 > 0 && dp[x1-1][x2-1] >= 0 {
					best = max(best, dp[x1-1][x2-1]+delta)
				}
				if x1 > 0 && y2 > 0 && dp[x1-1][x2] >= 0 {
					best = max(best, dp[x1-1][x2]+delta)
				}
				if y1 > 0 && x2 > 0 && dp[x1][x2-1] >= 0 {
					best = max(best, dp[x1][x2-1]+delta)
				}
				if y1 > 0 && y2 > 0 && dp[x1][x2] >= 0 {
					best = max(best, dp[x1][x2]+delta)
				}
				dp[x1][x2] = best
			}
		}
	}

	return max(dp[n-1][n-1], 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
