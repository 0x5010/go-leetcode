package leetcode0693

func hasAlternatingBits(n int) bool {
	n ^= n >> 2
	return n&(n-1) == 0
}
