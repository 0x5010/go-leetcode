package leetcode0960

func minDeletionSize(A []string) int {
	m, n := len(A), len(A[0])
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	res := n - 1
	for j := 0; j < n; j++ {
	L:
		for i := 0; i < j; i++ {
			for k := 0; k < m; k++ {
				if A[k][i] > A[k][j] {
					continue L
				}
			}
			if dp[i]+1 > dp[j] {
				dp[j] = dp[i] + 1
			}
		}
		res = min(res, n-dp[j])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
