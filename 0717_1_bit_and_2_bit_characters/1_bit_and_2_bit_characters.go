package leetcode0717

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i := 0
	for i < n-1 {
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}
	return i == n-1
}
