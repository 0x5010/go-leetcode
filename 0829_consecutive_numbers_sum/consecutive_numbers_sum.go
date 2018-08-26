package leetcode0829

func consecutiveNumbersSum(N int) int {
	n := mySqrt(2 * N)
	res := 1
	for k := 2; k <= n; k++ {
		if (N-(k*(k-1)/2))%k == 0 {
			res++
		}
	}
	return res
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
