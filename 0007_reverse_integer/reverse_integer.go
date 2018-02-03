package leetcode0007

import (
	"math"
)

func reverse(x int) int {
	sign := 1

	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		tmp := x % 10
		res = res*10 + tmp
		x = x / 10
	}
	res = sign * res

	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}
	return res
}
