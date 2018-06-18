package leetcode0459

import "strings"

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	return strings.Contains((s + s)[1:n*2-1], s)
}
