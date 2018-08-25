package leetcode0823

import "sort"

func numFactoredBinaryTrees(A []int) int {
	mod := 1000000007
	n := len(A)
	sort.Ints(A)
	dp := make(map[int]int, n)
	for i, a := range A {
		dp[a] = 1
		for j := 0; j < i; j++ {
			if a%A[j] == 0 && dp[a/A[j]] != 0 {
				dp[a] += dp[A[j]] * dp[a/A[j]]
			}
		}
	}
	res := 0
	for _, v := range dp {
		res += v
	}
	return res % mod
}
