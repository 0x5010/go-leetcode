package leetcode0044

func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)

	dp := make([][]bool, sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]bool, pLen+1)
	}

	dp[0][0] = true
	for i := 1; i <= pLen; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if p[j-1] != '*' {
				dp[i][j] = (p[j-1] == s[i-1] || p[j-1] == '?') && dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[sLen][pLen]
}
