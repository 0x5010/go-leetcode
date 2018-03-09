package leetcode0131

func partition(s string) [][]string {
	n := len(s)
	res := [][]string{}
	var dfs func(int, []string)
	dfs = func(index int, l []string) {
		if index == n {
			tmp := make([]string, len(l))
			copy(tmp, l)
			res = append(res, tmp)
			return
		}

		for i := index; i < n; i++ {
			if check(s[index : i+1]) {
				dfs(i+1, append(l, s[index:i+1]))
			}
		}
	}
	dfs(0, make([]string, 0, n))
	return res
}

func check(s string) bool {
	if len(s) < 2 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
