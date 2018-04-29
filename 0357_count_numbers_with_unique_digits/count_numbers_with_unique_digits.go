package leetcode0357

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n > 10 {
		n = 10
	}
	dp := make([]int, n)
	dp[0] = 10
	num, accum := 9, 9
	for i := 1; i < n; i++ {
		accum *= num
		num--
		dp[i] += accum + dp[i-1]
	}
	return dp[n-1]
}
