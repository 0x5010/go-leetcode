package leetcode0790

func numTilings(N int) int {
	if N < 3 {
		return N
	} else if N == 3 {
		return 5
	}
	mod := 1000000007
	dp := make([]int, N+1)
	dp[1], dp[2], dp[3] = 1, 2, 5
	for i := 4; i <= N; i++ {
		dp[i] = 2*dp[i-1] + dp[i-3]
		dp[i] %= mod
	}
	return dp[N]
}
