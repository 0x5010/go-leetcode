package leetcode0054

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}

	n := len(matrix[0])
	if n == 0 {
		return nil
	}

	if m == 1 {
		return matrix[0]
	}

	res := make([]int, 0, m*n)

	res = append(res, matrix[0]...)
	for i := 1; i < m-1; i++ {
		res = append(res, matrix[i][n-1])
	}
	for i := n - 1; i >= 0; i-- {
		res = append(res, matrix[m-1][i])
	}
	for i := m - 2; i > 0 && n > 1; i-- {
		res = append(res, matrix[i][0])
	}

	if m == 2 || n <= 2 {
		return res
	}

	next := make([][]int, m-2)
	for i := 0; i < m-2; i++ {
		next[i] = matrix[i+1][1 : n-1]
	}
	return append(res, spiralOrder(next)...)
}
