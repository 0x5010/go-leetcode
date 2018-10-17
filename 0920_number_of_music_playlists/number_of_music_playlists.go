package leetcode0920

func numMusicPlaylists(N int, L int, K int) int {
	mod := 1000000007
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, L+1)
	}
	dp[1][1] = 1
	for i := 1; i < N+1; i++ {
		for j := 1; j < L+1; j++ {
			if i == 1 && j == 1 {
				continue
			}
			v := (dp[i-1][j-1] * i) % mod
			if i-K > 0 {
				v = (v + (dp[i][j-1]*(i-K))%mod) % mod
			}
			dp[i][j] = v
		}
	}
	return dp[N][L]
}
