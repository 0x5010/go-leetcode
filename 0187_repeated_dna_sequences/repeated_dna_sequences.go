package leetcode0187

func findRepeatedDnaSequences(s string) []string {
	m := map[string]int{}
	for i := 0; i <= len(s)-10; i++ {
		sub := s[i : i+10]
		m[sub]++
	}
	res := []string{}
	for sub, count := range m {
		if count > 1 {
			res = append(res, sub)
		}
	}
	return res
}
