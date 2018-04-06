package leetcode0301

func removeInvalidParentheses(s string) []string {
	res := []string{}
	var dfs func(string, int, int, [2]byte)
	dfs = func(s string, i, j int, par [2]byte) {
		c := 0
		for x := i; x < len(s); x++ {
			if s[x] == par[0] {
				c++
			} else if s[x] == par[1] {
				c--
			}
			if c >= 0 {
				continue
			}
			for y := j; y <= x; y++ {
				if s[y] == par[1] && (y == j || s[y-1] != par[1]) {
					dfs(s[:y]+s[y+1:], x, y, par)
				}
			}
			return
		}
		r := reverse(s)
		if par[0] == '(' {
			dfs(r, 0, 0, [2]byte{')', '('})
		} else {
			res = append(res, r)
		}
	}
	dfs(s, 0, 0, [2]byte{'(', ')'})
	return res
}

func reverse(s string) string {
	i, j := 0, len(s)-1
	bs := []byte(s)
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
