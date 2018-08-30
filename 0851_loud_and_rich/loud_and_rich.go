package leetcode0851

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	rs := make([][]int, n)
	for _, r := range richer {
		rs[r[1]] = append(rs[r[1]], r[0])
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if res[i] >= 0 {
			return res[i]
		}
		res[i] = i
		for _, j := range rs[i] {
			if quiet[res[i]] > quiet[dfs(j)] {
				res[i] = res[j]
			}
		}
		return res[i]
	}
	for i := 0; i < n; i++ {
		dfs(i)
	}
	return res
}
