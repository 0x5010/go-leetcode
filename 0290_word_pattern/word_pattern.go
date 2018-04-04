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
	for i := 0; i < len(pattern); i++ {
		c := pattern[i]
		sub, ok := m[c]
		if ok {
			if sub != words[i] {
				return false
			}
		} else {
			for _, v := range m {
				if v == words[i] {
					return false
				}
			}
			m[c] = words[i]
		}
	}
	return true
}
