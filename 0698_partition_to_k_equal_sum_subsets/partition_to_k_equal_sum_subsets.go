package leetcode0698

import (
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	sum, max := nums[0], nums[0]
	for i := 1; i < n; i++ {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	avg := sum / k
	if sum%k != 0 || avg < max {
		return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	used := make([]bool, n)
	var dfs func(int, int, int) bool
	dfs = func(k, start, cur int) bool {
		if k == 1 {
			return true
		}
		if cur == avg {
			return dfs(k-1, 0, 0)
		}
		for i := start; i < n; i++ {
			if !used[i] && cur+nums[i] <= avg {
				used[i] = true
				if dfs(k, i+1, cur+nums[i]) {
					return true
				}
				used[i] = false
			}
		}
		return false
	}
	return dfs(k, 0, 0)
}
