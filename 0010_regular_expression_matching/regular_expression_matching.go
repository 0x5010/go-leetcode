package leetcode00010

func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	dp := make([][]bool, sLen+1)

	for i := range dp {
		dp[i] = make([]bool, pLen+1)
	}

	dp[sLen][pLen] = true
	for i := sLen; i > -1; i-- {
		for j := pLen - 1; j > -1; j-- {

			firstMatch := i < sLen && (p[j] == s[i] || p[j] == '.')
			if j+1 < pLen && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || firstMatch && dp[i+1][j]
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}
