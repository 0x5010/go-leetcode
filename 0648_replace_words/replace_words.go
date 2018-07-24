package leetcode0648

import (
	"sort"
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	sort.Slice(dict, func(i, j int) bool {
		return len(dict[i]) > len(dict[j])
	})
	ss := strings.Split(sentence, " ")
	for i, s := range ss {
		for _, d := range dict {
			if strings.HasPrefix(s, d) {
				ss[i] = d
			}
		}
	}
	return strings.Join(ss, " ")
}
