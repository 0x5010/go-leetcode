package leetcode0820

func minimumLengthEncoding(words []string) int {
	m := map[string]struct{}{}
	for _, word := range words {
		m[word] = struct{}{}
	}
	for _, word := range words {
		for i := 1; i < len(word); i++ {
			delete(m, word[i:])
		}
	}
	res := 0
	for word := range m {
		res += len(word) + 1
	}
	return res
}
