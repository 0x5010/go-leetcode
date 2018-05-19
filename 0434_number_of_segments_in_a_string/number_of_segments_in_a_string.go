package leetcode0434

import "strings"

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}

	return len(strings.Fields(s))
}
