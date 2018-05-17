package leetcode0424

func characterReplacement(s string, k int) int {
	n := len(s)
	counts := [26]int{}
	maxLen, l, r := 0, 0, 0
	for ; r < n; r++ {
		counts[s[r]-'A']++
		maxLen = max(maxLen, counts[s[r]-'A'])
		if r-l+1-maxLen > k {
			counts[s[l]-'A']--
			l++
		}
	}
	return n - l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
