package leetcode0792

import (
	"sort"
	"strings"
)

func numMatchingSubseq(S string, words []string) int {
	isSub := func(w string) int {
		f := 0
		for i := 0; i < len(w); i++ {
			if index := strings.IndexByte(S[f:], w[i]); index >= 0 {
				f += index + 1
			} else {
				return 0
			}
		}
		return 1
	}
	sort.Strings(words)
	p, a := "", 1
	res := 0
	for _, word := range words {
		if word != p {
			p, a = word, isSub(word)
		}
		res += a
	}
	return res
}
