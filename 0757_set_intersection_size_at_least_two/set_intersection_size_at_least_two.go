package leetcode0757

import "sort"

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	l, r := intervals[0][1]-1, intervals[0][1]
	res := 2
	for _, interval := range intervals {
		a, b := interval[0], interval[1]
		if l < a && a <= r {
			res++
			l, r = r, b
		} else if r < a {
			res += 2
			l, r = b-1, b
		}
	}
	return res
}
