package leetcode0368

import "sort"

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	sort.Ints(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	next := make([]int, n)
	max, index := 1, 0

	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[j]%nums[i] != 0 {
				continue
			}
			if dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				next[i] = j
			}
			if dp[i] > max {
				max = dp[i]
				index = i
			}
		}

	}

	res := make([]int, max)
	for i := 0; i < max; i++ {
		res[i] = nums[index]
		index = next[index]
	}
	return res
}
