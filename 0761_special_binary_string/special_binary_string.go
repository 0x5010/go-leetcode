package leetcode0761

import (
	"sort"
	"strings"
)

func makeLargestSpecial(S string) string {
	n := len(S)
	if n == 0 || n == 2 {
		return S
	}

	ss := []string{}
	count, i := 0, 0
	for j := 0; j < n; j++ {
		if S[j] == '0' {
			count--
		} else {
			count++
		}

		if count == 0 {
			ss = append(ss, "1"+makeLargestSpecial(S[i+1:j])+"0")
			i = j + 1
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	return strings.Join(ss, "")
}
