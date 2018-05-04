package leetcode0383

func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	for _, b := range magazine {
		m[b-'a']++
	}

	for _, b := range ransomNote {
		m[b-'a']--
		if m[b-'a'] < 0 {
			return false
		}
	}
	return true
}
