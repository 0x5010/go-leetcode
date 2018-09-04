package leetcode0866

import (
	"math"
	"strconv"
)

func primePalindrome(N int) int {
	if N <= 2 {
		return 2
	} else if N == 3 {
		return 3
	} else if N <= 5 {
		return 5
	} else if N <= 7 {
		return 7
	} else if N <= 11 {
		return 11
	}

	for lm := genLM(N); lm <= 100000000; lm++ {
		p := genPalindrome(lm)
		if p >= N && isPrime(p) {
			return p
		}
	}
	return -1
}

func genLM(N int) int {
	n := len(strconv.Itoa(N))
	base := int(math.Pow10(n / 2))
	if n%2 == 0 {
		return base
	}
	return N / base
}

func genPalindrome(lm int) int {
	for i := lm / 10; i > 0; i /= 10 {
		lm = lm*10 + i%10
	}
	return lm
}

func isPrime(x int) bool {
	if x <= 3 {
		return x > 1
	}
	if x%2 == 0 || x%3 == 0 {
		return false
	}
	for i := 5; i < mySqrt(x)+1; i += 6 {
		if x%i == 0 || x%(i+2) == 0 {
			return false
		}
	}
	return true
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
