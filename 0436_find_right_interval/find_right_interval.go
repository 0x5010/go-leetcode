package leetcode0436

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

import "sort"

func findRightInterval(intervals []Interval) []int {
	n := len(intervals)

	starts := make([]int, n)
	indexs := make(map[int]int, n)
	res := make([]int, n)

	for i, interval := range intervals {
		starts[i] = interval.Start
		indexs[interval.Start] = i
	}

	sort.Ints(starts)
	for i, interval := range intervals {
		index := sort.SearchInts(starts, interval.End)
		if index < n {
			res[i] = indexs[starts[index]]
		} else {
			res[i] = -1
		}
	}
	return res
}
