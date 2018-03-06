package leetcode0118

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	res := make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		for j := 1; j < i; j++ {
			tmp[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = tmp
	}
	return res
}
