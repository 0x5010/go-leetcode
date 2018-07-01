package leetcode0498

func findDiagonalOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])

	i, j, d := 0, 0, 0
	res := make([]int, m*n)
	dirs := [2][2]int{{-1, 1}, {1, -1}}
	for k := 0; k < m*n; k++ {
		res[k] = matrix[i][j]

		i += dirs[d][0]
		j += dirs[d][1]
		if i >= m {
			i = m - 1
			j += 2
			d = 1 - d
		}
		if j >= n {
			j = n - 1
			i += 2
			d = 1 - d
		}
		if i < 0 {
			i = 0
			d = 1 - d
		}
		if j < 0 {
			j = 0
			d = 1 - d
		}
	}
	return res
}
