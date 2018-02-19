package leetcode0059

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	top, bottom, left, right := 0, n-1, 0, n-1
	num := 1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++

		for i := top; i <= bottom; i++ {
			res[i][right] = num
			num++
		}
		right--

		for i := right; i >= left; i-- {
			res[bottom][i] = num
			num++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}
