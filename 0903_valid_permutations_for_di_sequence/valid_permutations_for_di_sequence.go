package leetcode0903

func numPermsDISequence(S string) int {
	mod := 1000000007
	n := len(S)
	dp0, dp1 := make([]int, n+1), make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp0[i] = 1
	}
	for i := 0; i < n; i++ {
		cur := 0
		if S[i] == 'I' {
			j := 0
			for ; j < n-i; j++ {
				cur = (cur + dp0[j]) % mod
				dp1[j] = cur
			}
		} else {
			j := n - i - 1
			for ; j >= 0; j-- {
				cur = (cur + dp0[j+1]) % mod
				dp1[j] = cur
			}
		}
		dp0, dp1 = dp1, dp0
	}
	return dp0[0]
}
