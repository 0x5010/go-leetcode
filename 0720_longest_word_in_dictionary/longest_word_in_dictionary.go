package leetcode0720

import "sort"

func longestWord(words []string) string {
	sort.Strings(words)
	m := map[string]struct{}{}
	res := words[0]
	for _, word := range words {
		n := len(word)
		if n == 1 {
			m[word] = struct{}{}
		} else if _, ok := m[word[:n-1]]; ok {
			m[word] = struct{}{}
			if len(word) > len(res) {
				res = word
			}
		}
	}
	return res
}
