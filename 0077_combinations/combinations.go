package leetcode0077

func combine(n int, k int) [][]int {
	c := make([]int, k)
	res := [][]int{}

	var dfs func(int, int)
	dfs = func(index, start int) {
		if index == k {
			tmp := make([]int, k)
			copy(tmp, c)
			res = append(res, tmp)
			return
		}

		for i := start; i <= n+1-k+index; i++ {
			c[index] = i
			dfs(index+1, i+1)
		}
	}
	dfs(0, 1)
	return res
}
