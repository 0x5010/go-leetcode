package leetcode0472

import "sort"

func findAllConcatenatedWordsInADict(words []string) []string {
	res := []string{}
	preWords := make(map[string]struct{})
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	canFrom := func(word string) bool {
		n := len(word)
		if len(preWords) == 0 {
			return false
		}
		dp := make([]bool, n+1)
		dp[0] = true
		for i := 1; i < n+1; i++ {
			for j := 0; j < i; j++ {
				if !dp[j] {
					continue
				}
				if _, ok := preWords[word[j:i]]; ok {
					dp[i] = true
					break
				}
			}
		}
		return dp[n]
	}

	for _, word := range words {
		if canFrom(word) {
			res = append(res, word)
		}
		preWords[word] = struct{}{}
	}
	return res
}
