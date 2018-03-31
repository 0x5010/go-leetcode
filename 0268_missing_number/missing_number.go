package leetcode0268

func missingNumber(nums []int) int {
	xor := 0
	for i, n := range nums {
		xor ^= i ^ n
	}
	return xor ^ len(nums)
}
