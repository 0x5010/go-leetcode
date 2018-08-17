package leetcode0807

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	r, c := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r[i] = max(r[i], grid[i][j])
			c[j] = max(c[j], grid[i][j])
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += min(r[i], c[j]) - grid[i][j]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
