package leetcode0566

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m := len(nums)
	if m == 0 {
		return nil
	}
	n := len(nums[0])
	if n == 0 {
		return nil
	}
	if m*n != r*c || (m == r && n == c) {
		return nums
	}

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)

		for j := 0; j < c; j++ {
			x := (i*c + j) / n
			y := (i*c + j) - x*n
			res[i][j] = nums[x][y]
		}
	}
	return res
}
