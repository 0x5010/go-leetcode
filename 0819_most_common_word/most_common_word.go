package leetcode0819

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	reg := regexp.MustCompile("[^a-z ]+")
	paragraph = reg.ReplaceAllString(strings.ToLower(paragraph), "")
	ss := strings.Fields(paragraph)
	m := map[string]struct{}{}
	for _, b := range banned {
		m[b] = struct{}{}
	}
	count := map[string]int{}
	max := 0
	res := ""
	for _, s := range ss {
		if _, ok := m[s]; ok {
			continue
		}
		count[s]++
		if count[s] > max {
			max = count[s]
			res = s
		}
	}
	return res
}
