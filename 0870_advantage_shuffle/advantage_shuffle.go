package leetcode0870

import (
	"sort"
)

func advantageCount(A []int, B []int) []int {
	n := len(A)
	bindex := make([][2]int, n)
	for i, b := range B {
		bindex[i][0], bindex[i][1] = b, i
	}
	sort.Slice(bindex, func(i, j int) bool {
		return bindex[i][0] < bindex[j][0]
	})
	sort.Ints(A)
	l, r := 0, n-1
	res := make([]int, n)
	for _, a := range A {
		if bindex[l][0] < a {
			res[bindex[l][1]] = a
			l++
		} else {
			res[bindex[r][1]] = a
			r--
		}
	}
	return res
}
