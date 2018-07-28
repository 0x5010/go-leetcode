package leetcode0688

import "math"

func knightProbability(N int, K int, r int, c int) float64 {
	dirs := [8][2]int{{1, 2}, {1, -2}, {2, 1}, {2, -1}, {-1, 2}, {-1, -2}, {-2, 1}, {-2, -1}}
	dp0 := make([]float64, N*N)
	dp1 := make([]float64, N*N)
	dp0[r*N+c] = 1

	for k := 0; k < K; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if dp0[i*N+j] == 0 {
					continue
				}
				for _, dir := range dirs {
					x, y := i+dir[0], j+dir[1]
					if x >= 0 && x < N && y >= 0 && y < N {
						dp1[x*N+y] += dp0[i*N+j]
					}
				}
			}
		}
		dp0, dp1 = dp1, dp0
		for i := range dp1 {
			dp1[i] = 0
		}
	}
	var count float64
	for _, v := range dp0 {
		count += v
	}
	return count / math.Pow(8, float64(K))
}
