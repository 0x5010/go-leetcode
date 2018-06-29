package leetcode0494

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < S {
		return 0
	}
	if (sum+S)%2 == 1 {
		return 0
	}

	target := (sum + S) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}
