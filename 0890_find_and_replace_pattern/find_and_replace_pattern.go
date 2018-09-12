package leetcode0890

func findAndReplacePattern(words []string, pattern string) []string {
	e := encode(pattern)
	res := []string{}
	for _, word := range words {
		if encode(word) == e {
			res = append(res, word)
		}
	}
	return res
}

func encode(s string) string {
	n := len(s)
	m := map[byte]byte{}
	for i := 0; i < n; i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = byte(len(m)) + 'a'
		}
	}
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = m[s[i]]
	}
	return string(res)
}
