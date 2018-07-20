package leetcode0633

func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}
	l, r := 0, mySqrt(c)
	for l <= r {
		cur := l*l + r*r
		if cur < c {
			l++
		} else if cur > c {
			r--
		} else {
			return true
		}
	}
	return false
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	low := 1
	high := x
	for low < high {
		mid := (low + high) / 2
		sq := mid * mid
		if sq > x {
			high = mid
		} else if sq < x {
			low = mid + 1
		} else {
			return mid
		}
	}

	return high - 1
}
