package leetcode0959

func regionsBySlashes(grid []string) int {
	n := len(grid)
	n3 := n * 3
	g := make([][]bool, n3)
	for i := 0; i < n3; i++ {
		g[i] = make([]bool, n3)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '/' {
				g[i*3][j*3+2], g[i*3+1][j*3+1], g[i*3+2][j*3] = true, true, true
			}
			if grid[i][j] == '\\' {
				g[i*3][j*3], g[i*3+1][j*3+1], g[i*3+2][j*3+2] = true, true, true
			}
		}
	}

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i >= 0 && j >= 0 && i < n3 && j < n3 && !g[i][j] {
			g[i][j] = true
			dfs(i-1, j)
			dfs(i+1, j)
			dfs(i, j-1)
			dfs(i, j+1)
		}
	}
	res := 0
	for i := 0; i < n3; i++ {
		for j := 0; j < n3; j++ {
			if !g[i][j] {
				dfs(i, j)
				res++
			}
		}
	}
	return res
}
