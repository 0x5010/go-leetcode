package leetcode0073

func setZeroes(matrix [][]int) {
	r := make([]bool, len(matrix))
	c := make([]bool, len(matrix[0]))

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				r[i] = true
				c[j] = true
			}
		}
	}

	for i := range r {
		if r[i] {
			for j := range matrix[i] {
				matrix[i][j] = 0
			}
		}
	}

	for j := range c {
		if c[j] {
			for i := range matrix {
				matrix[i][j] = 0
			}
		}
	}
}
