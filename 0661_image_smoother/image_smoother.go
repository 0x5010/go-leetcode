package leetcode0661

func imageSmoother(M [][]int) [][]int {
	m, n := len(M), len(M[0])
	res := make([][]int, m)

	average := func(r, c int) int {
		value, count := 0, 0
		for i := r - 1; i < r+2; i++ {
			for j := c - 1; j < c+2; j++ {
				if 0 <= i && i < m && 0 <= j && j < n {
					value += M[i][j]
					count++
				}
			}
		}
		return value / count
	}

	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = average(i, j)
		}
	}
	return res
}
