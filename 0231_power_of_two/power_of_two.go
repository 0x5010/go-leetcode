package leetcode0231

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	return n&(n-1) == 0
}
