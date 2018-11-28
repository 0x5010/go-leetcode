package leetcode0945

import "sort"

func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	need := 0
	res := 0
	for _, a := range A {
		res += max(need-a, 0)
		need = max(a, need) + 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
