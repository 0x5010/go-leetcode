package leetcode0764

func orderOfLargestPlusSign(N int, mines [][]int) int {
	grid := make([]int, N*N)
	for i := 0; i < N*N; i++ {
		grid[i] = N
	}
	for _, mine := range mines {
		grid[mine[0]*N+mine[1]] = 0
	}
	for i := 0; i < N; i++ {
		l, r, u, d := 0, 0, 0, 0
		for j, k := 0, N-1; j < N; j, k = j+1, k-1 {
			if grid[i*N+j] == 0 {
				l = 0
			} else {
				l++
			}
			if grid[i*N+k] == 0 {
				r = 0
			} else {
				r++
			}
			if grid[j*N+i] == 0 {
				u = 0
			} else {
				u++
			}
			if grid[k*N+i] == 0 {
				d = 0
			} else {
				d++
			}

			grid[i*N+j] = min(grid[i*N+j], l)
			grid[i*N+k] = min(grid[i*N+k], r)
			grid[j*N+i] = min(grid[j*N+i], u)
			grid[k*N+i] = min(grid[k*N+i], d)
		}
	}
	res := 0
	for i := 0; i < N*N; i++ {
		res = max(res, grid[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
