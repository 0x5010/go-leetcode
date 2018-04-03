package leetcode0279

func numSquares(n int) int {
	if isSquare(n) {
		return 1
	}

	for n%4 == 0 {
		n >>= 2
	}

	if n%8 == 7 {
		return 4
	}
	for i := 1; i <= mySqrt(n); i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}
	return 3
}

func isSquare(x int) bool {
	tmp := mySqrt(x)
	return x == tmp*tmp
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
