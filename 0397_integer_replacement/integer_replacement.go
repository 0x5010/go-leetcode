package leetcode0397

func integerReplacement(n int) int {
	res := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else if n == 3 || (n/2)%2 == 0 {
			n--
		} else {
			n++
		}
		res++
	}
	return res
}
