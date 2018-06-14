package leetcode463

func islandPerimeter(grid [][]int) int {
	islands, neighbours := 0, 0
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				islands++
				if i < m-1 && grid[i+1][j] == 1 {
					neighbours++
				}
				if j < n-1 && grid[i][j+1] == 1 {
					neighbours++
				}
			}
		}
	}
	return islands*4 - neighbours*2
}
