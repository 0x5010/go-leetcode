package leetcode0686

import "strings"

func repeatedStringMatch(A string, B string) int {
	var s strings.Builder
	repeat := len(B) / len(A)
	if len(B)%len(A) != 0 {
		repeat++
	}
	for i := 0; i < repeat; i++ {
		s.WriteString(A)
	}
	if strings.Index(s.String(), B) != -1 {
		return repeat
	}
	s.WriteString(A)
	if strings.Index(s.String(), B) != -1 {
		return repeat + 1
	}
	return -1
}
