package leetcode0514

import "math"

func findRotateSteps(ring string, key string) int {
	m, n := len(key), len(ring)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		if ring[j] == key[0] {
			dp[0][j] = min(j, n-j) + 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if ring[j] != key[i] {
				continue
			}

			for k := 0; k < n; k++ {
				if ring[k] == key[i-1] {
					step := dp[i-1][k] + min((j-k+n)%n, (k-j+n)%n) + 1
					if dp[i][j] == 0 {
						dp[i][j] = step
					} else {
						dp[i][j] = min(dp[i][j], step)
					}
				}
			}
		}
	}
	res := math.MaxInt32
	for j := 0; j < n; j++ {
		if dp[m-1][j] != 0 && dp[m-1][j] < res {
			res = dp[m-1][j]
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
