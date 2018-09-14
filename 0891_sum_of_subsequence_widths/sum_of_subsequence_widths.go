package leetcode0891

import "sort"

func sumSubseqWidths(A []int) int {
	mod := 1000000007
	sort.Ints(A)

	n, p := len(A), 1
	res := 0
	for i, a := range A {
		res = (res + (a-A[n-i-1])*p) % mod
		p = p * 2 % mod
	}
	return res
}
