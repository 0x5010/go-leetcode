package leetcode0387

func firstUniqChar(s string) int {
	m := make([]int, 26)
	for _, b := range s {
		m[b-'a']++
	}

	for i, b := range s {
		if m[b-'a'] == 1 {
			return i
		}
	}
	return -1
}
