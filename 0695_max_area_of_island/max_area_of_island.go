package leetcode0695

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dfs(i, j))
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
