package leetcode0289

func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])

	var lives func(int, int)
	lives = func(i, j int) {
		count := 0
		for x := i - 1; x <= i+1; x++ {
			for y := j - 1; y <= j+1; y++ {
				if x >= 0 && y >= 0 && x < m && y < n {
					count += board[x][y] % 2
				}
			}
		}
		count -= board[i][j] % 2

		if board[i][j] == 1 && (count == 2 || count == 3) {
			board[i][j] = 3
		} else if board[i][j] == 0 && count == 3 {
			board[i][j] = 2
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			lives(i, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] /= 2
		}
	}
}
