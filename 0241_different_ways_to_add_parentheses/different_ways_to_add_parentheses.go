package leetcode0241

import "strconv"

func diffWaysToCompute(input string) []int {
	cache := make(map[string][]int)
	var dfs func(string) []int
	dfs = func(s string) []int {
		if t, ok := cache[s]; ok {
			return t
		}
		tmp := []int{}
		for i := 0; i < len(s); i++ {
			if s[i] == '+' || s[i] == '-' || s[i] == '*' {
				for _, l := range dfs(s[:i]) {
					for _, r := range dfs(s[i+1:]) {
						tmp = append(tmp, operate(l, r, s[i]))
					}
				}
			}
		}

		if len(tmp) == 0 {
			num, _ := strconv.Atoi(s)
			tmp = append(tmp, num)
		}

		cache[s] = tmp
		return tmp
	}
	return dfs(input)
}

func operate(a, b int, opt byte) int {
	if opt == '+' {
		return a + b
	} else if opt == '-' {
		return a - b
	} else {
		return a * b
	}
}
