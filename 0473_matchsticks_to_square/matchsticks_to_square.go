package leetcode0473

import (
	"sort"
)

func makesquare(nums []int) bool {
	n := len(nums)
	if n < 4 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%4 != 0 {
		return false
	}
	avg := sum / 4

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	edges := [4]int{}

	var dfs func(int) bool
	dfs = func(index int) bool {
		if index == n {
			if edges[0] == avg && edges[1] == avg && edges[2] == avg {
				return true
			}
			return false
		}
		for i := 0; i < 4; i++ {
			if edges[i]+nums[index] > avg {
				continue
			}
			j := i
			for j >= 0 {
				j--
				if edges[i] == edges[j] {
					break
				}
			}
			if j != -1 {
				continue
			}
			edges[i] += nums[index]
			if dfs(index + 1) {
				return true
			}
			edges[i] -= nums[index]
		}
		return false
	}
	return dfs(0)
}
