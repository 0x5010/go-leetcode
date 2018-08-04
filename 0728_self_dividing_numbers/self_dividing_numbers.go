package leetcode0728

func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for i := left; i <= right; i++ {
		j := i
		for j > 0 {
			r := j % 10
			if r == 0 || i%r != 0 {
				break
			}
			j /= 10
		}
		if j == 0 {
			res = append(res, i)
		}
	}
	return res
}
