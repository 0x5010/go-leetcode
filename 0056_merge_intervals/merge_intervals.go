package leetcode0056

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

import (
	"sort"
)

func merge(intervals []Interval) []Interval {
	n := len(intervals)
	if n == 0 {
		return []Interval{}
	} else if n == 1 {
		return intervals
	}
	sortByStart(intervals)

	res := []Interval{intervals[0]}

	for _, cur := range intervals[1:] {
		previous := res[len(res)-1]
		if cur.Start <= previous.End {
			end := previous.End
			if cur.End > end {
				end = cur.End
			}
			res = append(res[:len(res)-1], Interval{Start: previous.Start})
		} else {
			res = append(res, cur)
		}
	}
	return res
}

func sortByStart(intervals []Interval) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
