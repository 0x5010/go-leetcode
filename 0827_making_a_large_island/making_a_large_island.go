package leetcode0827

func largestIsland(grid [][]int) int {
	n := len(grid)
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	area := make([][]int, n)
	for i := 0; i < n; i++ {
		area[i] = make([]int, n)
	}

	m := map[int]int{}
	m[0] = 0
	index := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 && area[i][j] < 1 {
				cur := 0
				connect := make([]int, 0, n)
				connect = append(connect, i*n+j)
				area[i][j] = index
				for cur < len(connect) {
					r, c := connect[cur]/n, connect[cur]%n
					for _, dir := range dirs {
						x, y := r+dir[0], c+dir[1]
						if x >= 0 && x < n && y >= 0 && y < n &&
							grid[x][y] > 0 && area[x][y] < 1 {
							area[x][y] = index
							connect = append(connect, x*n+y)
						}
					}
					cur++
				}
				m[index] = len(connect)
				index++
			}
		}
	}

	res := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := m[area[i][j]]
			if sum == 0 {
				sum = 1
				flag := map[int]bool{}
				for _, dir := range dirs {
					x, y := i+dir[0], j+dir[1]
					if x >= 0 && x < n && y >= 0 && y < n {
						flag[area[x][y]] = true
					}
				}
				for index := range flag {
					sum += m[index]
				}
			}
			if res < sum {
				res = sum
			}
		}
	}
	return res
}
