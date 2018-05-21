package leetcode0435

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

import "sort"

func eraseOverlapIntervals(intervals []Interval) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sortByEnd(intervals)
	end := intervals[0].End
	res := 0

	for i := 1; i < n; i++ {
		if intervals[i].Start >= end {
			end = intervals[i].End
		} else {
			res++
		}
	}
	return res
}

func sortByEnd(intervals []Interval) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].End < intervals[j].End
	})
}
