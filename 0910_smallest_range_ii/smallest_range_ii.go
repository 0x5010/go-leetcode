package leetcode0910

import "sort"

func smallestRangeII(A []int, K int) int {
	n := len(A)
	sort.Ints(A)
	res := A[n-1] - A[0]
	p, q := A[n-1]-K, A[0]+K
	for i := 0; i < n-1; i++ {
		res = min(res, max(p, A[i]+K)-min(q, A[i+1]-K))
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
