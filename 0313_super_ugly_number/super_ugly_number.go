package leetcode0313

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}

	pos := make([]int, len(primes))
	candidates := make([]int, len(primes))
	copy(candidates, primes)

	dp := make([]int, n)

	dp[0] = 1

	for i := 1; i < n; i++ {
		dp[i] = min(candidates)
		for j := 0; j < len(primes); j++ {
			if dp[i] == candidates[j] {
				pos[j]++
				candidates[j] = dp[pos[j]] * primes[j]
			}
		}
	}

	return dp[n-1]
}

func min(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}
