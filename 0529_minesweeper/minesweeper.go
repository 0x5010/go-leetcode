package leetcode0529

func updateBoard(board [][]byte, click []int) [][]byte {
	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}

	m, n := len(board), len(board[0])

	var dfs func(int, int)
	dfs = func(x, y int) {
		count := 0
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n {
				continue
			}
			if board[r][c] == 'M' || board[r][c] == 'X' {
				count++
			}
		}

		if count > 0 {
			board[x][y] = byte('0' + count)
			return
		}

		board[x][y] = 'B'
		for _, d := range dirs {
			r, c := x+d[0], y+d[1]
			if r < 0 || r >= m || c < 0 || c >= n {
				continue
			}
			if board[r][c] == 'E' {
				dfs(r, c)
			}
		}
	}
	dfs(x, y)
	return board
}
