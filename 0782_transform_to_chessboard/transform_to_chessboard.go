package leetcode0782

func movesToChessboard(board [][]int) int {
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[0][0]^board[i][0]^board[0][j]^board[i][j] == 1 {
				return -1
			}
		}
	}
	rSum, cSum, rSwap, cSwap := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		rSum += board[0][i]
		cSum += board[i][0]
		if board[i][0] == i%2 {
			rSwap++
		}
		if board[0][i] == i%2 {
			cSwap++
		}
	}
	if rSum < n/2 || (n+1)/2 < rSum || cSum < n/2 || (n+1)/2 < cSum {
		return -1
	}
	if n%2 == 1 {
		if rSwap%2 == 1 {
			rSwap = n - rSwap
		}
		if cSwap%2 == 1 {
			cSwap = n - cSwap
		}
	} else {
		rSwap = min(n-rSwap, rSwap)
		cSwap = min(n-cSwap, cSwap)
	}
	return (rSwap + cSwap) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
