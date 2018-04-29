package leetcode0372

func superPow(a int, b []int) int {
	base := 1337

	powmod := func(a, k int) int {
		a %= base
		res := 1
		for i := 0; i < k; i++ {
			res = (res * a) % base
		}
		return res
	}

	n := len(b)
	if n == 0 {
		return 1
	}

	last := b[n-1]
	b = b[:n-1]
	return powmod(superPow(a, b), 10) * powmod(a, last) % base
}
