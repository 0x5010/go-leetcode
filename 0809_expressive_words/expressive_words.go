package leetcode0809

func expressiveWords(S string, words []string) int {
	n := len(S)
	res := 0
L:
	for _, word := range words {
		m := len(word)
		j := 0
		for i := 0; i < n; i++ {
			if j < m && S[i] == word[j] {
				j++
			} else if i > 1 && S[i] == S[i-1] && S[i-1] == S[i-2] {
			} else if 0 < i && i < n-1 && S[i-1] == S[i] && S[i] == S[i+1] {
			} else {
				continue L
			}
		}
		if j == m {
			res++
		}
	}
	return res
}
