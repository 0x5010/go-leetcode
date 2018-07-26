package leetcode0672

func flipLights(n int, m int) int {
	if m == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		if m == 1 {
			return 3
		}
		return 4
	}
	if m == 1 {
		return 4
	}
	if m == 2 {
		return 7
	}
	return 8
}
