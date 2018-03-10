package leetcode0140

func wordBreak(s string, wordDict []string) []string {
	m := make(map[string][]string)
	wordSet := make(map[string]struct{}, len(wordDict))
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	var recur func(string) []string
	recur = func(s string) []string {
		if v, ok := m[s]; ok {
			return v
		}
		res := []string{}
		if _, ok := wordSet[s]; ok {
			res = append(res, s)
		}
		for i := 1; i < len(s); i++ {
			cur := s[i:]
			if _, ok := wordSet[cur]; ok {
				strs := recur(s[:i])
				for _, sub := range strs {
					res = append(res, sub+" "+cur)
				}
			}
		}
		m[s] = res
		return res
	}
	return recur(s)
}
