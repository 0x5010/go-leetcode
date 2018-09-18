package leetcode0906

import (
	"math"
	"strconv"
)

func superpalindromesInRange(L string, R string) int {
	l, _ := strconv.Atoi(L)
	r, _ := strconv.Atoi(R)

	left := int(math.Floor(math.Sqrt(float64(l))))
	right := int(math.Ceil(math.Sqrt(float64(r))))

	res := 0
	for i := left; i < right; {
		p := nextP(i)
		if p <= right && isP(p*p) {
			res++
		}
		i = p + 1
	}
	return res
}

func nextP(i int) int {
	s := strconv.Itoa(i)
	n := len(s)
	h := s[:(n+1)/2]
	tmp, _ := strconv.Atoi(h)
	nh := strconv.Itoa(tmp + 1)
	p1, _ := strconv.Atoi(h[:n/2] + reverseString(h[:n/2]))
	p2, _ := strconv.Atoi(nh[:n/2] + reverseString(nh[:n/2]))
	cands := []int{int(math.Pow10(n) - 1), p1, p2}

	p := math.MaxInt32
	for _, c := range cands {
		if c >= i && i < p {
			p = i
		}
	}
	return p
}

func isP(i int) bool {
	s := strconv.Itoa(i)
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func reverseString(s string) string {
	bs, i, j := []byte(s), 0, len(s)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
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
