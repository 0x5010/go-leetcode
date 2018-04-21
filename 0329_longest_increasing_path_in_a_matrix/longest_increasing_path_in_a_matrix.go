package leetcode0329

import (
	"math"
)

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	path := make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}

	res := 0

	var dfs func(int, int, int) int
	dfs = func(i, j, pre int) int {

		if i < 0 || i >= m || j < 0 || j >= n || matrix[i][j] >= pre {
			return 0
		}

		if path[i][j] > 0 {
			return path[i][j]
		}
		tmp := 0
		tmp = max(dfs(i-1, j, matrix[i][j]), tmp)
		tmp = max(dfs(i+1, j, matrix[i][j]), tmp)
		tmp = max(dfs(i, j-1, matrix[i][j]), tmp)
		tmp = max(dfs(i, j+1, matrix[i][j]), tmp)
		tmp++
		path[i][j] = tmp
		return tmp
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(dfs(i, j, math.MaxInt32), res)
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
