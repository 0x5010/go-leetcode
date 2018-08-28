package leetcode0837

func new21Game(N int, K int, W int) float64 {
	if K == 0 || N >= K+W {
		return 1.
	}
	dp := make([]float64, N+1)
	dp[0] = 1.
	s, m := 1., 1/float64(W)
	res := 0.
	for i := 1; i <= N; i++ {
		dp[i] = s * m
		if i < K {
			s += dp[i]
		} else {
			res += dp[i]
		}
		if i >= W {
			s -= dp[i-W]
		}
	}
	return res
}
