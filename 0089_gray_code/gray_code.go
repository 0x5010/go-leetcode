package leetcode0089

func grayCode(n int) []int {
	count := 1 << uint32(n)
	res := make([]int, count)
	for i := 0; i < count; i++ {
		res[i] = i ^ i>>1
	}
	return res
}
