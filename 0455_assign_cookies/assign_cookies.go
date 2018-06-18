package leetcode0455

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	n := len(g)
	sort.Ints(g)
	sort.Ints(s)

	res := 0
	for _, x := range s {
		if res >= n {
			break
		}
		if x >= g[res] {
			res++
		}
	}
	return res
}
