package leetcode0840

func numMagicSquaresInside(grid [][]int) int {
	dirs := [][2]int{
		{-1, -1}, {1, 1}, {-1, 1}, {1, -1}, // corners
		{0, -1}, {0, 1}, {-1, 0}, {1, 0}, // edges
	}
	isCorner := func(a int) bool {
		switch a {
		case 2:
		case 4:
		case 6:
		case 8:
		default:
			return false
		}
		return true
	}
	isMagic := func(i, j int) bool {
		if grid[i][j] != 5 {
			return false
		}
		pre := 0
		for index, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			v := grid[x][y]
			if index < 4 {
				if !isCorner(v) {
					return false
				}
			} else if v < 1 || v > 9 {
				return false
			}
			if index%2 == 0 {
				pre = v
			} else {
				if pre+v != 10 {
					return false
				}
			}
		}
		return true
	}
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if isMagic(i, j) {
				res++
			}
		}
	}
	return res
}
