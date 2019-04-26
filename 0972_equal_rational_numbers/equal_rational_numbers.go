package leetcode0972

import (
	"strconv"
	"strings"
)

func isRationalEqual(S string, T string) bool {
	return compute(S) == compute(T)
}

func compute(s string) float64 {
	i := strings.Index(s, "(")
	if i != -1 {
		base := s[:i]
		rep := s[i+1:]
		for j := 0; j < 20; j++ {
			base += rep
		}
		s = base
	}
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
