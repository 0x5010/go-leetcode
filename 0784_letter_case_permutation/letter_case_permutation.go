package leetcode0784

import "strings"

func letterCasePermutation(S string) []string {
	res := []string{}
	n := len(S)
	var recur func(string, int)
	recur = func(cur string, i int) {
		if i == n {
			res = append(res, cur)
			return
		}
		if '0' <= S[i] && S[i] <= '9' {
			recur(cur+S[i:i+1], i+1)
		} else {
			recur(cur+strings.ToLower(S[i:i+1]), i+1)
			recur(cur+strings.ToUpper(S[i:i+1]), i+1)
		}
	}
	recur("", 0)
	return res
}
