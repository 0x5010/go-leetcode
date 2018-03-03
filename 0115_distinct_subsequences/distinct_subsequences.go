package leetcode00115

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[j+1] += dp[j]
			}
		}
	}
	return dp[n]
}
