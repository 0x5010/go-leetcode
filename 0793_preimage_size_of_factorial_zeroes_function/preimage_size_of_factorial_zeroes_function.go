package leetcode0793

func preimageSizeFZF(K int) int {
	l, r := 0, 5*(K+1)
	for l <= r {
		m := (l + r) / 2
		zeros := 0
		x := m
		for x > 0 {
			x /= 5
			zeros += x
		}
		if zeros == K {
			return 5
		} else if zeros < K {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return 0
}
