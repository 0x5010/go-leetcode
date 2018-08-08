package leetcode0749

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func containVirus(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for {
		okMap := make([][]bool, m)
		for i := 0; i < m; i++ {
			okMap[i] = make([]bool, n)
		}
		blk := make([][]int, 0)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 1 && okMap[i][j] == false {
					blk = append(blk, search(grid, i, j, m, n, okMap))
				}
			}
		}
		chosen := []int{}
		best := 0
		walls := 0
		for i := 0; i < len(blk); i++ {
			if blk[i][2] > best {
				best = blk[i][2]
				chosen = blk[i]
				walls = blk[i][3]
			}
		}
		if best == 0 {
			return res
		}
		res += walls
		clear(grid, chosen[0], chosen[1], m, n)

		for i := 0; i < len(blk); i++ {
			if blk[i][0] != chosen[0] || blk[i][1] != chosen[1] {
				infect(grid, blk[i][0], blk[i][1], m, n)
			}
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 2 {
					grid[i][j] = 1
				}
			}
		}
	}
}

func search(grid [][]int, i, j, m, n int, ok [][]bool) []int {
	ok[i][j] = true
	res := []int{i, j}
	queue := [][2]int{{i, j}}
	num := 0
	effected := make([][]bool, m)
	for i := 0; i < len(effected); i++ {
		effected[i] = make([]bool, n)
	}

	needWall := 0
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			if grid[nx][ny] == 1 && ok[nx][ny] == false {
				ok[nx][ny] = true
				queue = append(queue, [2]int{nx, ny})
			} else if effected[nx][ny] == false && grid[nx][ny] == 0 {
				effected[nx][ny] = true
				num++
				needWall++
			} else if grid[nx][ny] == 0 {
				needWall++
			}
		}
	}
	res = append(res, num)
	res = append(res, needWall)
	return res
}

func clear(grid [][]int, i, j, m, n int) {
	queue := [][2]int{{i, j}}
	used := make([]bool, m*n)
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		used[x*n+y] = true
		queue = queue[1:]
		grid[x][y] = 3
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			if grid[nx][ny] == 1 && !used[nx*n+ny] {
				queue = append(queue, [2]int{nx, ny})
				used[nx*n+ny] = true
			}
		}
	}
}

func infect(grid [][]int, i, j, m, n int) {
	queue := [][2]int{{i, j}}
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		used[x][y] = true
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			if grid[nx][ny] == 1 && used[nx][ny] == false {
				queue = append(queue, [2]int{nx, ny})
				used[nx][ny] = true
			} else if grid[nx][ny] == 0 {
				grid[nx][ny] = 2
			}
		}
	}
}
