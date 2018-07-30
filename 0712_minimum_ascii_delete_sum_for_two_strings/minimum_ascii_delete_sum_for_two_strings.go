package leetcode0712

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp0 := make([]int, n+1)
	dp1 := make([]int, n+1)
	dp0[0] = 0
	for i := 1; i <= n; i++ {
		dp0[i] = dp0[i-1] + int(s2[i-1])
	}
	for i := 1; i <= m; i++ {
		dp1[0] = dp0[0] + int(s1[i-1])
		for j := 1; j <= n; j++ {
			both := 0
			if s1[i-1] != s2[j-1] {
				both += int(s1[i-1])
				both += int(s2[j-1])
			}
			dp1[j] = min(
				dp0[j-1]+both,
				min(dp0[j]+int(s1[i-1]), dp1[j-1]+int(s2[j-1])),
			)
		}
		dp0, dp1 = dp1, dp0
	}
	return dp0[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
