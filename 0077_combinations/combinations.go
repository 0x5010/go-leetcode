package leetcode0077

func combine(n int, k int) [][]int {
	comb := make([]int, k)
	res := make([][]int, 0, c(n, k))

	var dfs func(int, int)
	dfs = func(index, start int) {
		if index == k {
			tmp := make([]int, k)
			copy(tmp, comb)
			res = append(res, tmp)
			return
		}

		for i := start; i <= n+1-k+index; i++ {
			comb[index] = i
			dfs(index+1, i+1)
		}
	}
	dfs(0, 1)
	return res
}

func c(n int, k int) int {
	res := 1
	for i := n; i > n-k; i-- {
		res *= i
	}
	for i := 1; i < k+1; i++ {
		res /= i
	}
	return res
}
