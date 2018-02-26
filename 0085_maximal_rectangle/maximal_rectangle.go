package leetcode0085

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	left := make([]int, n)
	right := make([]int, n)
	height := make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	max := 0

	for i := 0; i < m; i++ {
		l, r := 0, n
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[j]++

				if l > left[j] {
					left[j] = l
				}
			} else {
				height[j] = 0

				left[j] = 0
				l = j + 1
			}
		}
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if r < right[j] {
					right[j] = r
				}
			} else {
				right[j] = n
				r = j
			}
		}

		for j := 0; j < n; j++ {
			tmp := (right[j] - left[j]) * height[j]
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}
