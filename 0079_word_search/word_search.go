package leetcode0079

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}

	n := len(board[0])
	if n == 0 {
		return false
	}

	wLen := len(word)
	if wLen == 0 {
		return false
	}

	var dfs func(int, int, int) bool
	dfs = func(r, c, i int) bool {
		if wLen == i {
			return true
		}

		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != word[i] {
			return false
		}

		tmp := board[r][c]
		board[r][c] = 0
		i++
		if dfs(r-1, c, i) || dfs(r+1, c, i) || dfs(r, c-1, i) || dfs(r, c+1, i) {
			return true
		}
		board[r][c] = tmp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
