package leetcode0526

func countArrangement(N int) int {
	if N < 2 {
		return N
	}

	used := make([]bool, N+1)

	var dfs func(int) int
	dfs = func(n int) int {
		if n == 1 {
			return 1
		}
		res, l := 0, len(used)
		for i := n; i < l; i += n {
			if used[i] == false {
				used[i] = true
				res += dfs(n - 1)
				used[i] = false
			}
		}
		for i := 2; i <= l; i++ {
			t := n / i
			if n%i == 0 && used[t] == false {
				used[t] = true
				res += dfs(n - 1)
				used[t] = false
			}
		}
		return res
	}
	return dfs(N)
}
