package leetcode0539

import (
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	m := map[int]struct{}{}
	l := make([]int, 0, len(timePoints))
	for _, timePoint := range timePoints {
		i := Ttoi(timePoint)
		if _, ok := m[i]; ok {
			return 0
		}
		m[i] = struct{}{}
		l = append(l, i)
	}
	sort.Ints(l)

	n := len(l)
	res := l[0] - l[n-1] + 24*60
	for i := 1; i < n; i++ {
		res = min(res, l[i]-l[i-1])
	}
	return res
}

func Ttoi(s string) int {
	h, _ := strconv.Atoi(s[:2])
	m, _ := strconv.Atoi(s[3:])
	return 60*h + m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
