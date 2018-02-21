package leetcode0065

import (
	"strings"
)

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	pointSeen, eSeen, numberSeen := false, false, false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			numberSeen = true
		} else if c == '.' {
			if eSeen || pointSeen {
				return false
			}
			pointSeen = true
		} else if c == 'e' {
			if eSeen || !numberSeen {
				return false
			}
			numberSeen = false
			eSeen = true
		} else if c == '-' || c == '+' {
			if i != 0 && s[i-1] != 'e' {
				return false
			}
		} else {
			return false
		}
	}
	return numberSeen
}
