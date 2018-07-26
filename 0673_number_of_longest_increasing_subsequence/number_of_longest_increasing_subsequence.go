package leetcode0673

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	max := 0
	res := 0
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i][0], dp[i][1] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i][0] == dp[j][0]+1 {
					dp[i][1] += dp[j][1]
				} else if dp[i][0] < dp[j][0]+1 {
					dp[i][0], dp[i][1] = dp[j][0]+1, dp[j][1]
				}
			}
		}
		if max == dp[i][0] {
			res += dp[i][1]
		} else if max < dp[i][0] {
			max = dp[i][0]
			res = dp[i][1]
		}
	}
	return res
}
