package leetcode0811

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := map[string]int{}
	for _, c := range cpdomains {
		ss := strings.Fields(c)
		n, _ := strconv.Atoi(ss[0])
		s := ss[1]
		m[s] += n
		for i := 0; i < len(s); i++ {
			if s[i] == '.' {
				d := s[i+1:]
				m[d] += n
			}
		}
	}
	res := make([]string, 0, len(m))
	for d, n := range m {
		res = append(res, fmt.Sprintf("%d %s", n, d))
	}
	return res
}
