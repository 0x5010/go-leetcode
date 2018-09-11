package leetcode0887

func superEggDrop(K int, N int) int {
	dp := make([]int, K+1)
	res := 0
	for dp[K] < N {
		for i := K; i > 0; i-- {
			dp[i] += dp[i-1] + 1
		}
		res++
	}
	return res
}
