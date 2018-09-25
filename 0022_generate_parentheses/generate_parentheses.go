package leetcode0022

func generateParenthesis(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	dp[1] = []string{"()"}
	for i := 2; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			for _, l := range dp[j] {
				for _, r := range dp[i-j-1] {
					dp[i] = append(dp[i], "("+l+")"+r)
				}
			}
		}
	}
	return dp[n]
}
