package leetcode0477

func totalHammingDistance(nums []int) int {
	res := 0
	n := len(nums)
	for i := 0; i < 32; i++ {
		bitCount := 0
		for _, num := range nums {
			bitCount += (num >> uint(i)) & 1
		}
		res += bitCount * (n - bitCount)
	}
	return res
}
