package leetcode0312

func maxCoins(nums []int) int {
	n := len(nums) + 2
	inums := make([]int, n)
	inums[0], inums[n-1] = 1, 1
	copy(inums[1:], nums)

	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}

	for i := 2; i < n; i++ {
		for l := 0; l < n-i; l++ {
			r := l + i
			for j := l + 1; j < r; j++ {
				tmp := m[l][j] + inums[l]*inums[j]*inums[r] + m[j][r]
				if tmp > m[l][r] {
					m[l][r] = tmp
				}

			}
		}
	}
	return m[0][n-1]
}
