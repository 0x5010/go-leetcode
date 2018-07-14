package leetcode0576

func findPaths(m int, n int, N int, i int, j int) int {
	const mod = 1e9 + 7
	dp := [50][50]int{}

	for k := 0; k < N; k++ {
		pre := make([]int, n)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				paths := 0

				if i == 0 {
					paths++
				} else {
					paths += pre[j]
				}

				if j == 0 {
					paths++
				} else {
					paths += pre[j-1]
				}

				if i == m-1 {
					paths++
				} else {
					paths += dp[i+1][j]
				}

				if j == n-1 {
					paths++
				} else {
					paths += dp[i][j+1]
				}

				paths %= mod
				pre[j] = dp[i][j]
				dp[i][j] = paths
			}
		}
	}
	return dp[i][j]
}
