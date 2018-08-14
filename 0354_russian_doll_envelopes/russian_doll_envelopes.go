package leetcode0354

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n < 2 {
		return n
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := make([]int, n)
	res := 1
	for i, envelope := range envelopes {
		index := sort.Search(res, func(idx int) bool {
			return envelopes[dp[idx]][1] >= envelope[1]
		})
		dp[index] = i
		if index >= res {
			res++
		}
	}
	return res
}
