package leetcode0803

func hitBricks(grid [][]int, hits [][]int) []int {
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n, l := len(grid), len(grid[0]), len(hits)
	for _, h := range hits {
		grid[h[0]][h[1]]--
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n ||
			grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = 2
		count := 1
		for _, dir := range dirs {
			count += dfs(i+dir[0], j+dir[1])
		}
		return count
	}
	connected := func(i, j int) bool {
		if i == 0 {
			return true
		}
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n &&
				grid[x][y] == 2 {
				return true
			}
		}
		return false
	}
	for i := 0; i < n; i++ {
		dfs(0, i)
	}
	res := make([]int, l)
	for k := l - 1; k >= 0; k-- {
		i, j := hits[k][0], hits[k][1]
		grid[i][j]++
		if grid[i][j] == 1 && connected(i, j) {
			res[k] = dfs(i, j) - 1
		}
	}
	return res
}
