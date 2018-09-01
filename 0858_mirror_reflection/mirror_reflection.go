package leetcode0858

func mirrorReflection(p int, q int) int {
	l := p * q / gcd(p, q)
	if (l/q)%2 == 0 {
		return 2
	}
	return (l / p) % 2
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
