package leetcode0878

func nthMagicalNumber(N int, A int, B int) int {
	mod := 1000000007
	if A < B {
		A, B = B, A
	}
	lcm := A * B / gcd(A, B)
	l, r := B*N/2, A*N
	for l < r {
		m := (l + r) / 2
		if m/A+m/B-m/lcm < N {
			l = m + 1
		} else {
			r = m
		}
	}
	return l % mod
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
