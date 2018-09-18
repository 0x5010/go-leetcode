package leetcode0906

import (
	"math"
	"strconv"
)

func superpalindromesInRange(L string, R string) int {
	l, _ := strconv.ParseInt(L, 10, 0)
	r, _ := strconv.ParseInt(R, 10, 0)

	left := int(math.Sqrt(float64(l)))
	right := int(math.Sqrt(float64(r)))

	res := 0
L:
	for n := 1; n < 19; n++ {
		for _, p := range genP(n)[1:] {
			if p < left {
				continue
			}
			if p > right {
				break L
			}
			if isP(p * p) {
				res++
			}
		}
	}
	return res
}

func genP(n int) []int {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
	if n == 2 {
		return []int{0, 11, 22, 33, 44, 55, 66, 77, 88, 99}
	}
	sub := genP(n - 2)
	res := []int{}
	for i := 0; i < 10; i++ {
		for _, s := range sub {
			r := i
			for j := 0; j < n-1; j++ {
				r *= 10
			}
			res = append(res, r+s*10+i)
		}
	}
	return res
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
