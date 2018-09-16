package leetcode0902

import "strconv"

func atMostNGivenDigitSet(D []string, N int) int {
	n := strconv.Itoa(N)
	ld, ln := len(D), len(n)
	res := 0
	for i := 1; i < ln; i++ {
		res += pow(ld, i)
	}
	for i := 0; i < ln; i++ {
		same := false
		for _, d := range D {
			if d[0] < n[i] {
				res += pow(ld, ln-i-1)
			} else if d[0] == n[i] {
				same = true
			}
		}
		if !same {
			return res
		}
	}
	return res + 1
}

func pow(x, n int) int {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	res := pow(x, n>>1)
	if n&1 == 0 {
		return res * res
	}
	return res * res * x
}
