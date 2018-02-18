package leetcode0057

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

func insert(intervals []Interval, newInterval Interval) []Interval {
	s, e := newInterval.Start, newInterval.End
	l, r := []Interval{}, []Interval{}
	for _, i := range intervals {
		if i.End < s {
			l = append(l, i)
		} else if i.Start > e {
			r = append(r, i)
		} else {
			if i.Start < s {
				s = i.Start
			}
			if i.End > e {
				e = i.End
			}
		}
	}
	l = append(l, Interval{Start: s, End: e})
	return append(l, r...)
}
