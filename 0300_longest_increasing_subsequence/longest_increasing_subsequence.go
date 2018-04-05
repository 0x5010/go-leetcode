package leetcode0300

import (
	"sort"
)

func lengthOfLIS(nums []int) int {
	tail, res := make([]int, len(nums)), 0

	for _, n := range nums {
		i := sort.SearchInts(tail[:res], n)
		tail[i] = n
		if i == res {
			res++
		}
	}
	return res
}
