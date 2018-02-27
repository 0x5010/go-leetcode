package leetcode0097

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([]bool, n+1)
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[0] = true
			} else if i == 0 {
				dp[j] = dp[j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[j] = dp[j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) ||
					(dp[j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return dp[n]
}
