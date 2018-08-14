package leetcode0788

func rotatedDigits(N int) int {
	res := 0
	for i := 2; i <= N; i++ {
		if valid(i) {
			res++
		}
	}
	return res
}

func valid(n int) bool {
	r := false
	for n > 0 {
		switch n % 10 {
		case 2, 5, 6, 9:
			r = true
		case 3, 4, 7:
			return false
		}
		n /= 10
	}
	return r
}
