package leetcode0475

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	res := 0

	sort.Ints(houses)
	sort.Ints(heaters)

	m, n := len(houses), len(heaters)

	index := sort.SearchInts(heaters, houses[0])
	for _, house := range houses {
		for index < n && house > heaters[index] {
			index++
		}

		if index == n {
			return max(res, houses[m-1]-heaters[index-1])
		}
		left := math.MaxInt32
		if index > 0 {
			left = house - heaters[index-1]
		}
		right := heaters[index] - house
		res = max(res, min(left, right))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
