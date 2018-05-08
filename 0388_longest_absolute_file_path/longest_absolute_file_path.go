package leetcode0388

import "strings"

func lengthLongestPath(input string) int {
	paths := strings.Split(input, "\n")
	stack := make([]int, len(paths)+1)
	res := 0
	for _, s := range paths {
		lev := strings.LastIndex(s, "\t") + 1
		stack[lev+1] = stack[lev] + len(s) - lev + 1
		if strings.Contains(s, ".") {
			if stack[lev+1]-1 > res {
				res = stack[lev+1] - 1
			}
		}
	}
	return res
}
