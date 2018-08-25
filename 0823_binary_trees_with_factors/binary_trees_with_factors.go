package leetcode0823

import "sort"

func numFactoredBinaryTrees(A []int) int {
	mod := 1000000007
	n := len(A)
	sort.Ints(A)
	dp := make([]int, n)
	dp[0] = 1
	res := 1
	for i := 1; i < n; i++ {
		c := 1
		for j, k := 0, i-1; j <= k; {
			jk := A[j] * A[k]
			if jk == A[i] {
				p := dp[j] * dp[k]
				if j != k {
					p += p
				}
				c += p
				j++
				k--
			} else if jk < A[i] {
				j++
			} else {
				k--
			}
		}
		dp[i] = c
		res += c
	}
	return res % mod
}
