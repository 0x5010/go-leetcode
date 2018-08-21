package leetcode0813

func largestSumOfAverages(A []int, K int) float64 {
	n := len(A)
	prev := make([]float64, n)
	sum := 0.
	for i, a := range A {
		sum += float64(a)
		prev[i] = sum / float64(i+1)
	}
	dp := make([]float64, n)
	for i := 2; i <= K; i++ {
		for j := i; j <= n; j++ {
			s, m := 0., 0.
			for k := j - 1; k >= i-1; k-- {
				s += float64(A[k])
				v := s/float64(j-k) + prev[k-1]
				if v > m {
					m = v
				}
			}
			dp[j-1] = m
		}
		prev, dp = dp, prev
	}
	return prev[n-1]
}
