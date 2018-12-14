package leetcode0948

import "sort"

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	i, j, p := 0, len(tokens)-1, 0
	res := 0
	for i <= j {
		if P >= tokens[i] {
			P -= tokens[i]
			i++
			p++
			res = max(res, p)
		} else if p > 0 {
			p--
			P += tokens[j]
			j--
		} else {
			break
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
