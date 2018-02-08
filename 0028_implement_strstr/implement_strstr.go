package leetcode0028

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	for i := 0; i < len(haystack)-n+1; i++ {
		j := 0
		for j < n {
			if haystack[i+j] != needle[j] {
				break
			}
			j++
		}
		if j == n {
			return i
		}
	}
	return -1
}
