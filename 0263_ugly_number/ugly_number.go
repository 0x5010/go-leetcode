package leetcode0263

func isUgly(num int) bool {
	if num == 0 {
		return false
	}

	for _, n := range []int{2, 3, 5} {
		for num%n == 0 {
			num /= n
		}
	}
	return num == 1
}
