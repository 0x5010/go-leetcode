package leetcode0935

func knightDialer(N int) int {
	mod := 1000000007
	dp := make([]int, 10)
	for i := 0; i < 10; i++ {
		dp[i] = 1
	}
	f := func(n ...int) int {
		sum := 0
		for _, i := range n {
			sum = (sum + i) % mod
		}
		return sum
	}
	for i := 0; i < N-1; i++ {
		dp[1], dp[2], dp[3],
			dp[4], dp[5], dp[6],
			dp[7], dp[8], dp[9],
			dp[0] =
			f(dp[6], dp[8]), f(dp[7], dp[9]), f(dp[4], dp[8]),
			f(dp[3], dp[9], dp[0]), 0, f(dp[1], dp[7], dp[0]),
			f(dp[2], dp[6]), f(dp[1], dp[3]), f(dp[2], dp[4]),
			f(dp[4], dp[6])
	}

	return f(dp...)
}
