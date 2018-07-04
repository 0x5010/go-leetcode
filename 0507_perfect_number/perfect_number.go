package leetcode0507

func checkPerfectNumber(num int) bool {
	if num%2 == 1 {
		return false
	}

	if num%10 != 6 && num%100 != 28 {
		return false
	}

	sum := 1
	for i := 2; i <= mySqrt(num); i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}
	return sum == num
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
