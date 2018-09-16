package leetcode0899

import "sort"

func orderlyQueue(S string, K int) string {
	if K > 1 {
		bs := []byte(S)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		return string(bs)
	}
	n := len(S)
	res := S
	for i := 1; i < n; i++ {
		tmp := S[i:] + S[:i]
		if res > tmp {
			res = tmp
		}
	}
	return res
}
