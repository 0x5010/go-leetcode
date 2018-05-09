package leetcode0392

func isSubsequence(s string, t string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	for i, j := 0, 0; j < len(t); j++ {
		if s[i] == t[j] {
			i++
			if i == n {
				return true
			}
		}
	}
	return false
}
