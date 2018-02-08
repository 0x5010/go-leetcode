package leetcode0029

import "math"

func d(dividend, divisor, count int) (int, int) {
	switch {
	case dividend < divisor:
		return 0, dividend
	case dividend >= divisor && dividend < divisor+divisor:
		return count, dividend - divisor
	default:
		quotient, remainder := d(dividend, divisor+divisor, count+count)
		if remainder >= divisor {
			return quotient + count, remainder - divisor
		}
		return quotient, remainder
	}
}

func iabs(num int) (abs, sign int) {
	if num < 0 {
		return 0 - num, -1
	}
	return num, 1
}

func divide(dividend int, divisor int) int {
	if divisor == 0 || (dividend == math.MinInt32 && divisor == -1) {
		return math.MaxInt32
	}
	m, signM := iabs(dividend)
	n, signN := iabs(divisor)
	sign := 1
	if signM != signN {
		sign = -1
	}

	quotient, _ := d(m, n, 1)
	if sign == -1 {
		return 0 - quotient
	}
	return quotient
}
