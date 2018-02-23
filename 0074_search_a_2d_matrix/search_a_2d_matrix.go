package leetcode0074

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	r := 0
	for r < m && matrix[r][0] <= target {
		r++
	}
	r--

	i, j := 0, n-1
	for i <= j {
		middle := (i + j) / 2
		if matrix[r][middle] == target {
			return true
		} else if matrix[r][middle] > target {
			j = middle - 1
		} else {
			i = middle + 1
		}
	}
	return false
}
