package leetcode0646

import "sort"

func findLongestChain(pairs [][]int) int {
	n := len(pairs)
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] > pairs[j][0]
		}
		return pairs[i][1] < pairs[j][1]
	})

	res := 1
	b := pairs[0][1]
	for i := 1; i < n; i++ {
		c := pairs[i][0]
		if b < c {
			res++
			b = pairs[i][1]
		}
	}
	return res
}
