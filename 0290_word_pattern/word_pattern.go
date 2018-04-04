package leetcode0290

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	if len(words) != len(pattern) {
		return false
	}
	m := make(map[byte]string)
	seen := make(map[string]struct{})
	for i := 0; i < len(pattern); i++ {
		c, s := pattern[i], words[i]
		v, ok := m[c]
		if ok {
			if v != s {
				return false
			}
		} else {
			if _, ok := seen[s]; ok {
				return false
			}
			m[c] = s
			seen[s] = struct{}{}
		}
	}
	return true
}
