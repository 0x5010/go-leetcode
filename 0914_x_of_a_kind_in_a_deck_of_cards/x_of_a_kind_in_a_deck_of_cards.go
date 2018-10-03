package leetcode0914

func hasGroupsSizeX(deck []int) bool {
	count := map[int]int{}
	x := 0
	for _, d := range deck {
		count[d]++
	}
	for _, c := range count {
		x = gcd(c, x)
	}
	return x > 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
