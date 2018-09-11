package leetcode0884

import "strings"

func uncommonFromSentences(A string, B string) []string {
	m := map[string]int{}
	for _, s := range strings.Fields(A) {
		m[s]++
	}
	for _, s := range strings.Fields(B) {
		m[s]++
	}
	res := []string{}
	for s, c := range m {
		if c == 1 {
			res = append(res, s)
		}
	}
	return res
}
