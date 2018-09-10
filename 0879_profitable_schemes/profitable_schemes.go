package leetcode0879

func profitableSchemes(G int, P int, group []int, profit []int) int {
	mod := 1000000007
	n := len(group)
	dp := make([][]int, P+1)
	for i := 0; i < P+1; i++ {
		dp[i] = make([]int, G+1)
	}
	dp[0][0] = 1
	for k := 0; k < n; k++ {
		for i := P; i >= 0; i-- {
			ip := min(i+profit[k], P)
			for j := G - group[k]; j >= 0; j-- {
				dp[ip][j+group[k]] += dp[i][j]
				dp[ip][j+group[k]] %= mod
			}
		}
	}
	res := 0
	for i := 0; i < G+1; i++ {
		res += dp[P][i]
	}
	return res % mod
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
