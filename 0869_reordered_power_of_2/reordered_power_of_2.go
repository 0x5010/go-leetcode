package leetcode0869

func reorderedPowerOf2(N int) bool {
	e := encode(N)
	for i := uint(0); i < 32; i++ {
		if encode(1<<i) == e {
			return true
		}
	}
	return false
}

func encode(x int) int {
	e := 0
	n := uint(x)
	for n > 0 {
		d := n % 10
		e += 1 << (4 * d)
		n /= 10
	}
	return e
}
