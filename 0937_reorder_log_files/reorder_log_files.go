package leetcode0937

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	n := len(logs)
	letters, digits := make([]string, 0, n), []string{}
	for _, log := range logs {
		if log[strings.IndexByte(log, ' ')+1] <= '9' {
			digits = append(digits, log)
		} else {
			letters = append(letters, log)
		}
	}

	sort.Slice(letters, func(i, j int) bool {
		s1, s2 := letters[i], letters[j]
		return s1[strings.IndexByte(s1, ' ')+1:] < s2[strings.IndexByte(s2, ' ')+1:]
	})
	for _, log := range digits {
		letters = append(letters, log)
	}
	return letters
}
