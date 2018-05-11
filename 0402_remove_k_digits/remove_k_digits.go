package leetcode0402

func removeKdigits(num string, k int) string {
	n := len(num)
	digits := n - k
	stack := make([]byte, n)
	top := 0

	for i := 0; i < n; i++ {
		for top > 0 && stack[top-1] > num[i] && k > 0 {
			top--
			k--
		}
		stack[top] = num[i]
		top++
	}

	i := 0
	for i < digits && stack[i] == '0' {
		i++
	}
	if i == digits {
		return "0"
	}
	return string(stack[i:digits])
}
