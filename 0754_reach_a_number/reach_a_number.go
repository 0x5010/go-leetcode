package leetcode0754

import "math"

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	n := int(math.Ceil((math.Sqrt(8*float64(target)+1) - 1) / 2))
	sum := n * (n + 1) / 2
	if sum == target {
		return n
	}
	x := sum - target
	if x&1 == 0 {
		return n
	}
	if n&1 == 1 {
		return n + 2
	}
	return n + 1
}
