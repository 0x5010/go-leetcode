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
	ll, lr := len(strconv.Itoa(left)), len(strconv.Itoa(right))

	res := 0
L:
	for n := ll; n < lr+1; n++ {
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

var cacheP = [18][]int{nil, {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {0, 11, 22, 33, 44, 55, 66, 77, 88, 99}}

func genP(n int) []int {
	if n == 0 {
		return nil
	}
	res := cacheP[n]
	if len(res) != 0 {
		return cacheP[n]
	}

	sub := cacheP[n-2]
	if len(sub) == 0 {
		sub = genP(n - 2)
	}
	res = []int{}
	for i := 0; i < 10; i++ {
		for _, s := range sub {
			r := i
			for j := 0; j < n-1; j++ {
				r *= 10
			}
			res = append(res, r+s*10+i)
		}
	}
	cacheP[n] = res
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
