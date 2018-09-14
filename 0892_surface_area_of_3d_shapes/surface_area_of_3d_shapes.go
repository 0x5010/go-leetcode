package leetcode0892

func surfaceArea(grid [][]int) int {
	n := len(grid)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := grid[i][j]
			if v == 0 {
				continue
			}
			res += 4*v + 2
			if i != 0 {
				res -= min(v, grid[i-1][j]) * 2
			}
			if j != 0 {
				res -= min(v, grid[i][j-1]) * 2
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
