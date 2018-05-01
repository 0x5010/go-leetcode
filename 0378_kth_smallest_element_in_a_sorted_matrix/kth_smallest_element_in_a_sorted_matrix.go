package leetcode0378

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)

	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		m := (l + r) / 2
		count := 0
		for i := 0; i < n; i++ {
			j := n - 1
			for j >= 0 && matrix[i][j] > m {
				j--
			}
			count += j + 1
		}

		if count < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
