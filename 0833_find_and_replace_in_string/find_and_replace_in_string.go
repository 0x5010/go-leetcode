package leetcode0833

import (
	"sort"
	"strings"
)

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	n, m := len(indexes), len(S)
	es := make([]*entry, n)
	for i, index := range indexes {
		es[i] = &entry{
			i: index,
			s: sources[i],
			t: targets[i],
		}
	}
	sort.Slice(es, func(i, j int) bool {
		return es[i].i < es[j].i
	})
	bs := strings.Builder{}
	i, j := 0, 0
	for _, e := range es {
		i = j
		bs.WriteString(S[i:e.i])
		j = min(m, e.i+len(e.s))
		s := S[e.i:j]
		if s == e.s {
			bs.WriteString(e.t)
		} else {
			bs.WriteString(s)
		}
	}
	if j < m {
		bs.WriteString(S[j:])
	}
	return bs.String()
}

type entry struct {
	i    int
	s, t string
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
