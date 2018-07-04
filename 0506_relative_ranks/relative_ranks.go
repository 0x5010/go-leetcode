package leetcode0506

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	n := len(nums)
	m := make(map[int]int, n)
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)
	for i, num := range tmp {
		m[num] = n - i
	}

	res := make([]string, n)
	for i, num := range nums {
		switch m[num] {
		case 1:
			res[i] = "Gold Medal"
		case 2:
			res[i] = "Silver Medal"
		case 3:
			res[i] = "Bronze Medal"
		default:
			res[i] = strconv.Itoa(m[num])
		}
	}
	return res
}
