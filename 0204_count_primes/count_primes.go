package leetcode0204

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	notPrime := make([]bool, n)
	count := n / 2
	for i := 3; i < mySqrt(n)+1; i += 2 {
		if notPrime[i] {
			continue
		}

		for j := i * i; j < n; j += 2 * i {
			if !notPrime[j] {
				count--
				notPrime[j] = true
			}
		}
	}
	return count
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
