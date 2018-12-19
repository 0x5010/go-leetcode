package leetcode0953

func isAlienSorted(words []string, order string) bool {
	indexs := [26]int{}
	for i := 0; i < len(order); i++ {
		indexs[order[i]-'a'] = i
	}

	compare := func(s1, s2 string) int {
		m, n, cmp := len(s1), len(s2), 0
		for i := 0; i < n && i < m && cmp == 0; i++ {
			cmp = indexs[s1[i]-'a'] - indexs[s2[i]-'a']
		}
		if cmp == 0 {
			return m - n
		}
		return cmp
	}

	for i := 1; i < len(words); i++ {
		if compare(words[i-1], words[i]) > 0 {
			return false
		}
	}
	return true
}
