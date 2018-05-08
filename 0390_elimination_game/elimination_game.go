package leetcode0390

func lastRemaining(n int) int {
	left := true
	step, head := 1, 1
	for n > 1 {
		if left || n%2 == 1 {
			head = head + step
		}
		n /= 2
		step *= 2
		left = !left
	}
	return head
}
