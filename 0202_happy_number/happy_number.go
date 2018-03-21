package leetcode0202

func isHappy(n int) bool {
	slow, fast := n, squares(n)
	for slow != fast {
		slow = squares(slow)
		fast = squares(squares(fast))
	}
	if slow == 1 {
		return true
	}
	return false
}

func squares(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}
